/**
 * Created by shiyi on 2017/11/24.
 * Email: shiyi@fightcoder.com
 */

package judge

import (
	"os"
	"time"
)

type JudgeJava struct {
	JudgeBase
}

func (this *JudgeJava) Compile() Result {
	os.Rename("./code.java", "Main.java")
	return this.compile("javac", []string{"Main.java"}, 5*time.Second)
}

func (this *JudgeJava) Run(inputFile string, outputFile string) Result {
	result := this.run("java", []string{"Main"}, inputFile, outputFile, (time.Duration)(1+this.MemoryLimit)*time.Second)
	return result
}
