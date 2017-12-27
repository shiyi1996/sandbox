/**
 * Created by shiyi on 2017/11/24.
 * Email: shiyi@fightcoder.com
 */

package judge

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Judge interface {
	Compile() Result
	Run(inFileCase string, outFileCase string) Result
}

func newJudge(language string, codeFile string, timeLimit int64, memoryLimit int64, outputLimit int64) Judge {
	switch language {
	case "cpp":
		return &JudgeCpp{
			JudgeBase{
				CodeFile:    codeFile,
				TimeLimit:   timeLimit,
				MemoryLimit: memoryLimit,
				OutputLimit: outputLimit,
			},
		}
	default:
		panic("No such language")
	}
}

func DoJudge(judgeType string, language string, codeFile string, timeLimit int64, memoryLimit int64, outputLimit int64) {
	judge := newJudge(language, codeFile, timeLimit, memoryLimit, outputLimit)

	if judgeType == "default" {
		doJudge(judge)
	} else if judgeType == "contest" {
		doContestJudge(judge)
	} else if judgeType == "test" {
		doTestJudge(judge)
	} else if judgeType == "special" {
		doSpecialJudge(judge)
	}
}

func notify() {

}

func getCaseList(path string) []string {
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

func doJudge(judge Judge) {
	result := judge.Compile()
	if result.ResultCode != 0 {
		fmt.Printf("Compile Error :%#v\n", result)
		return
	}

	caseList := getCaseList(getCurrentPath() + "/case")
	for _, name := range caseList {
		result = judge.Run("case/"+name+".in", "case/"+name+".out")
	}
}

func doContestJudge(judge Judge) {

}

func doTestJudge(judge Judge) {

}

func doSpecialJudge(judge Judge) {

}

func getCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}
