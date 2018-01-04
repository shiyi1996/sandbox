/**
 * Created by shiyi on 2017/11/24.
 * Email: shiyi@fightcoder.com
 */

package judge

import (
	"time"
)

type JudgeCpp struct {
	JudgeBase
}

func (this *JudgeCpp) Compile() Result {
	return this.compile("g++", []string{"code.cpp", "-fmax-errors=200"}, 5*time.Second)
}

func (this *JudgeCpp) Run(inputFile string, outputFile string) Result {
	result := this.run("./a.out", []string{}, inputFile, outputFile, (time.Duration)(1+this.MemoryLimit)*time.Second)
	return result
}
