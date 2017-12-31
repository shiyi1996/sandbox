package judge

import (
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
