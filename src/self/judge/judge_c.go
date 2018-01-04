/**
 * Created by shiyi on 2017/11/24.
 * Email: shiyi@fightcoder.com
 */

package judge

import (
	"time"
)

type JudgeC struct {
	JudgeBase
}

func (this *JudgeC) Compile() Result {
	return this.compile("gcc", []string{"code.c"}, 5*time.Second)
}

func (this *JudgeC) Run(inputFile string, outputFile string) Result {
	result := this.run("./a.out", []string{}, inputFile, outputFile, 2*time.Second)
	return result
}
