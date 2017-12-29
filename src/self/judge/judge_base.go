/**
 * Created by shiyi on 2017/12/16.
 * Email: shiyi@fightcoder.com
 */

package judge

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type JudgeBase struct {
	TimeLimit   int64 //Second
	MemoryLimit int64 //KB
	OutputLimit int64 //KB

}

func (this *JudgeBase) compile(cmdName string, cmdArg []string, timeout time.Duration) Result {
	cmd := exec.Command(cmdName, cmdArg...)

	var output []byte
	errC := make(chan error, 1)
	go func() {
		var err error
		output, err = cmd.CombinedOutput()
		errC <- err
	}()

	select {
	case <-time.After(timeout):
		cmd.Process.Kill()
		fmt.Println("KILL")
		return Result{
			ResultCode: CompilationError,
			ResultDes:  "Compile Timeout",
		}
	case err := <-errC:
		if err != nil {
			fmt.Println("报错", err)
			return Result{
				ResultCode: CompilationError,
				ResultDes:  string(output),
			}
		}
	}
	//if _, err := os.Stat(outBinName); os.IsNotExist(err) {
	//	stderrData, err := ioutil.ReadAll(stderr)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	return Result{
	//		ResultCode: 1,
	//		ResultDes:  string(stderrData),
	//	}
	//}

	return Result{
		ResultCode: Normal,
		ResultDes:  "",
	}
}

func (this *JudgeBase) run(cmdName string, cmdArg []string, inputFile string, outputFile string, timeout time.Duration) Result {
	cmd := exec.Command(cmdName, cmdArg...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	cmd.Stdin = bufio.NewReader(input)

	output, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer output.Close()
	outputWriter := bufio.NewWriter(output)
	defer outputWriter.Flush()
	cmd.Stdout = outputWriter

	timeoutStopFlag := false

	time.AfterFunc(timeout, func() {
		cmd.Process.Kill()
		timeoutStopFlag = true
		fmt.Println("KILL")
	})

	fmt.Println(time.Now())

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	var usage syscall.Rusage
	var wStatus syscall.WaitStatus

	_, err = syscall.Wait4(cmd.Process.Pid, &wStatus, syscall.WUNTRACED, &usage)
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Now())

	fmt.Printf("||| %#v \n", usage)

	sig := wStatus.Signal()
	fmt.Println(sig)
	//if wStatus.CoreDump() {
	//	return Result{
	//		ResultCode: RuntimeError,
	//		ResultDes:  string("adsasd"),
	//	}
	//}
	if sig == syscall.SIGSEGV {
		return Result{
			ResultCode:    RuntimeError,
			ResultDes:     sig.String(),
			RunningTime:   -1,
			RunningMemory: -1,
		}
	}
	if sig == syscall.SIGXCPU || usage.Utime.Sec > this.TimeLimit || timeoutStopFlag {
		return Result{
			ResultCode:    TimeLimitExceeded,
			ResultDes:     "",
			RunningTime:   -1,
			RunningMemory: -1,
		}
	}
	if sig == syscall.SIGXFSZ {
		return Result{
			ResultCode:    OutputLimitExceeded,
			ResultDes:     sig.String(),
			RunningTime:   -1,
			RunningMemory: -1,
		}
	}
	if usage.Maxrss > this.MemoryLimit*1024 {
		stderrData, err := ioutil.ReadAll(stderr)
		if err != nil {
			panic(err)
		}
		return Result{
			ResultCode:    MemoryLimitExceeded,
			ResultDes:     string(stderrData),
			RunningTime:   -1,
			RunningMemory: -1,
		}
	}
	if wStatus.Exited() {
		if wStatus.ExitStatus() != 0 {
			stderrData, err := ioutil.ReadAll(stderr)
			if err != nil {
				panic(err)
			}
			return Result{
				ResultCode:    RuntimeError,
				ResultDes:     string(stderrData),
				RunningTime:   -1,
				RunningMemory: -1,
			}
		}
	}

	useTime := (usage.Stime.Sec+usage.Utime.Sec)*1000 + int64(usage.Stime.Usec+usage.Utime.Usec)/1000
	useMemory := usage.Maxrss / 1024

	return Result{
		ResultCode:    Normal,
		ResultDes:     "",
		RunningTime:   useTime,
		RunningMemory: useMemory,
	}
}
