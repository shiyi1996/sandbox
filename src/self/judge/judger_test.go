package judge

import (
	"fmt"
	"os"
	"testing"
)

func TestNotify(t *testing.T) {
	var judger Judger
	judger.notify(Result{})
}

var judger = Judger{
	Language:    "c++",
	TimeLimit:   1,
	MemoryLimit: 128000,
	OutputLimit: 500,
}

func TestDoJudge(t *testing.T) {
	os.Chdir("/Users/shiyi/project/fightcoder/sandbox/tmp")

	judger.doJudge()
}

func TestCompare(t *testing.T) {
	os.Chdir("/Users/shiyi/project/fightcoder/sandbox/tmp")

	var judger Judger
	result := judger.compare("/Users/shiyi/project/fightcoder/sandbox/tmp/a", "/Users/shiyi/project/fightcoder/sandbox/tmp/b")
	fmt.Printf("%#v", result)
}
