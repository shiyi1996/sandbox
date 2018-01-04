/**
 * Created by shiyi on 2017/11/24.
 * Email: shiyi@fightcoder.com
 */

package judge

import (
	"time"
)

type JudgePy struct {
	JudgeBase
}

func (this *JudgePy) Compile() Result {
	return Result{ResultCode: Normal}
}

func (this *JudgePy) Run(inputFile string, outputFile string) Result {
	result := this.run("python", []string{"code.py"}, inputFile, outputFile, 2*time.Second)
	return result
}
