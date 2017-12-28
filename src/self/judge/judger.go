package judge

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func (this *Judger) notify() {

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

func (this *Judger) doJudge() {
	judge := newJudge(this.Language, this.TimeLimit, this.MemoryLimit, this.OutputLimit)

	result := judge.Compile()
	if result.ResultCode != 0 {
		fmt.Printf("Compile Error :%#v\n", result)
		return
	}

	caseList := this.getCaseList(getCurrentPath() + "/case")
	for _, name := range caseList {
		result = judge.Run("case/"+name+".in", "case/"+name+".out")
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
