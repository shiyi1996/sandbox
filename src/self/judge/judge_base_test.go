/**
 * Created by shiyi on 2017/12/18.
 * Email: shiyi@fightcoder.com
 */

package judge

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var judgeBase = JudgeBase{
	TimeLimit:   1,
	MemoryLimit: 12800000000,
	OutputLimit: 102456456,
}

func TestJudgeBasecompileSucc(t *testing.T) {
	os.Chdir("/Users/shiyi/project/fightcoder-sandbox/case")

	result := judgeBase.compile("g++", []string{"1.cpp"}, 2*time.Second)
	fmt.Println(result)
}

func TestJudgeBasecompileWarn(t *testing.T) {
	os.Chdir("/Users/shiyi/project/fightcoder-sandbox/case")

	result := judgeBase.compile("g++", []string{"2.cpp"}, 2*time.Second)
	fmt.Println(result)
}

func TestJudgeBasecompileError(t *testing.T) {
	os.Chdir("/Users/shiyi/project/fightcoder-sandbox/case")

	result := judgeBase.compile("g++", []string{"3.cpp"}, 2*time.Second)
	fmt.Println(result)
}

func TestJudgeBasecompileTimeout(t *testing.T) {
	os.Chdir("/Users/shiyi/project/fightcoder-sandbox/case")

	result := judgeBase.compile("g++", []string{"4.cpp"}, 2*time.Second)
	fmt.Println(result)
}

func TestJudgeBasecompileLog(t *testing.T) {
	os.Chdir("/Users/shiyi/project/fightcoder-sandbox/case")

	result := judgeBase.compile("g++", []string{"5.cpp"}, 2*time.Second)
	fmt.Println(result)
}

func TestJudgeBaseRun(t *testing.T) {
	os.Chdir("/Users/shiyi/project/fightcoder-sandbox/case")
	result := judgeBase.compile("g++", []string{"1.cpp"}, 5*time.Second)
	fmt.Println(result)

	result = judgeBase.run("/Users/shiyi/project/fightcoder-sandbox/case/a.out",
		[]string{}, "1.in", "1.out", 2*time.Second)
	fmt.Println(result)
	fmt.Println("time:", judgeBase.RunningTime)
	fmt.Println("memory:", judgeBase.RunningMemory)
}
