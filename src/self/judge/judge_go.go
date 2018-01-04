/**
 * Created by shiyi on 2017/11/24.
 * Email: shiyi@fightcoder.com
 */

package judge

import (
	"time"
)

type JudgeGo struct {
	JudgeBase
}

func (this *JudgeGo) Compile() Result {
	return this.compile("go", []string{"build", "code.go"}, 5*time.Second)
}

func (this *JudgeGo) Run(inputFile string, outputFile string) Result {
	result := this.run("./code", []string{}, inputFile, outputFile, (time.Duration)(1+this.MemoryLimit)*time.Second)
	return result
}
