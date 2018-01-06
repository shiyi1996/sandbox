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
	changeSubmitUrl = "http://128.0.9.207:8888/change_submit"
	//changeSubmitUrl = "http://128.0.9.207:8888/change_submit"
	workDir = "/workspace"
	//workDir = "/Users/shiyi/project/fightcoder/sandbox/tmp"
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
	fmt.Printf("%s\n", result.ResultDes)

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
	fmt.Printf("notify res: %s\n", r)
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
	cmd := exec.Command("diff", "-B", userOutput, caseOutput)

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
		RunningTime:   -1,
	})

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
		RunningTime:   -1,
	})

	totalResult := Result{
		ResultCode:    Accepted,
		ResultDes:     "",
		RunningMemory: 0,
		RunningTime:   0,
	}

	//caseList := this.getCaseList(getCurrentPath() + "/case")
	caseList := this.getCaseList(workDir + "/case")

	for _, name := range caseList {
		result = judge.Run(workDir+"/case/"+name+".in", workDir+"/output.txt")
		if result.ResultCode != Normal {
			fmt.Printf("Running Error :%#v\n", result)
			this.notify(result)
			return
		}

		if result.RunningMemory > totalResult.RunningMemory {
			totalResult.RunningMemory = result.RunningMemory
		}

		if result.RunningTime > totalResult.RunningTime {
			totalResult.RunningTime = result.RunningTime
		}

		result = this.compare(workDir+"/output.txt", workDir+"/case/"+name+".out")
		if result.ResultCode != Accepted {
			this.notify(result)
			return
		}
	}

	this.notify(totalResult)
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
