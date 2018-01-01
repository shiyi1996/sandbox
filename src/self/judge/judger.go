package judge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	changeSubmitUrl = "http://judgeip:8888/change_submit"
	//changeSubmitUrl = "http://128.0.9.207:8888/change_submit"
)

type Judger struct {
	SubmitType  string
	SubmitId    int64
	JudgeType   string
	Language    string
	TimeLimit   int64
	MemoryLimit int64
	OutputLimit int64
}

func (this *Judger) DoJudge() {
	if this.JudgeType == "default" {
		this.doJudge()
	} else if this.JudgeType == "contest" {
		this.doContestJudge()
	} else if this.JudgeType == "test" {
		this.doTestJudge()
	} else if this.JudgeType == "special" {
		this.doSpecialJudge()
	}
}

type ChangeSubMess struct {
	SubmitType string
	SubmitId   int64
	Result     Result
}

func (this *Judger) notify(result Result) {
	fmt.Printf("notify: %#v\n", result)

	changeSubMess := ChangeSubMess{
		Result:     result,
		SubmitType: this.SubmitType,
		SubmitId:   this.SubmitId,
	}

	data, err := json.Marshal(changeSubMess)
	if err != nil {
		panic(err)
	}

	body := bytes.NewBuffer([]byte(data))
	res, err := http.Post(changeSubmitUrl, "application/json;charset=utf-8", body)
	if err != nil {
		panic(err)
	}

	r, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}
	fmt.Printf("%s", r)
}

func (this *Judger) getCaseList(path string) []string {
	dir_list, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	caseList := make([]string, 0)

	for _, v := range dir_list {
		if v.IsDir() {
			continue
		}

		name := v.Name()
		if name[len(name)-3:] == ".in" {
			caseList = append(caseList, name[:len(name)-3])
		}
	}

	return caseList
}

func (this *Judger) compare(userOutput string, caseOutput string) Result {
	cmd := exec.Command("diff", "-B", "-b", userOutput, caseOutput)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return Result{
			ResultCode: WrongAnswer,
			ResultDes:  string(output),
		}
	}

	return Result{
		ResultCode: Accepted,
	}
}

func (this *Judger) doJudge() {
	judge := newJudge(this.Language, this.TimeLimit, this.MemoryLimit, this.OutputLimit)

	this.notify(Result{
		ResultCode:    Compiling,
		ResultDes:     "",
		RunningMemory: -1,
		RunningTime:   -1})

	result := judge.Compile()
	if result.ResultCode != 0 {
		fmt.Printf("Compile Error :%#v\n", result)
		this.notify(result)
		return
	}

	this.notify(Result{
		ResultCode:    Running,
		ResultDes:     "",
		RunningMemory: -1,
		RunningTime:   -1})

	//caseList := this.getCaseList("case")
	caseList := this.getCaseList(getCurrentPath() + "/case")
	for _, name := range caseList {
		result = judge.Run("case/"+name+".in", "output.txt")

		if result.ResultCode != 0 {
			fmt.Printf("Running Error :%#v\n", result)
			this.notify(result)
			return
		}

		//result = this.compare("output.txt", "case/"+name+".out")
		//this.notify(result)
	}
}

func (this *Judger) doContestJudge() {}

func (this *Judger) doTestJudge() {}

func (this *Judger) doSpecialJudge() {}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}
