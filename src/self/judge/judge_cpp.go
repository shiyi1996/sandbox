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
	return this.compile("g++", []string{"code.cpp"}, 5*time.Second)
}

func (this *JudgeCpp) Run(inFileCase string, outFileCase string) Result {
	result := this.run("./a.out", []string{}, inFileCase, "output.txt", 2*time.Second)
	//this.compare("output.txt", outFileCase)
	return result
}
