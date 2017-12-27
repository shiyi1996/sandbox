///**
// * Created by shiyi on 2017/12/16.
// * Email: shiyi@fightcoder.com
// */
//
package judge

//
///*
//	#include <unistd.h>
//
//	int my_fork() {
//		pid_t pid = fork();
//		return pid;
//	}
//*/
//import "C"
//import (
//	"bufio"
//	"fmt"
//	"io/ioutil"
//	"os"
//	"os/exec"
//	"syscall"
//	"time"
//)
//
//type JudgeBase struct {
//	CodeFile      string
//	TimeLimit     int64  //Second
//	MemoryLimit   int64  //KB
//	OutputLimit   int64  //KB
//	RunningTime   int64  //耗时(ms)
//	RunningMemory int64  //所占空间
//	Result        string //运行结果
//	ResultDes     string //结果描述
//}
//
//func (this *JudgeBase) Compile(cmdName string, cmdArg []string, timeout time.Duration) Result {
//	cmd := exec.Command(cmdName, cmdArg...)
//
//	var output []byte
//	errC := make(chan error, 1)
//	go func() {
//		var err error
//		output, err = cmd.CombinedOutput()
//		errC <- err
//	}()
//
//	select {
//	case <-time.After(timeout):
//		cmd.Process.Kill()
//		fmt.Println("KILL")
//		return Result{
//			ResultCode: CompilationError,
//			ResultDes:  "Compile Timeout",
//		}
//	case err := <-errC:
//		if err != nil {
//			fmt.Println("报错", err)
//			return Result{
//				ResultCode: CompilationError,
//				ResultDes:  string(output),
//			}
//		}
//	}
//	//if _, err := os.Stat(outBinName); os.IsNotExist(err) {
//	//	stderrData, err := ioutil.ReadAll(stderr)
//	//	if err != nil {
//	//		panic(err)
//	//	}
//	//
//	//	return Result{
//	//		ResultCode: 1,
//	//		ResultDes:  string(stderrData),
//	//	}
//	//}
//
//	return Result{
//		ResultCode: Normal,
//		ResultDes:  "",
//	}
//}
//
//func (this *JudgeBase) Run(cmdName string, cmdArg []string, inFile string, outputFile string, timeout time.Duration) Result {
//	pid, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
//	if err != 0 {
//		panic(err)
//	} else if pid > 0 {
//		fmt.Println("FATHER")
//		var timeLimitKilled = false
//
//		time.AfterFunc(timeout, func() {
//			syscall.Kill(int(pid), syscall.SIGKILL)
//			timeLimitKilled = true
//			fmt.Println("KILL KILL KILL")
//		})
//
//		var usage syscall.Rusage
//		var wStatus syscall.WaitStatus
//		if _, err := syscall.Wait4(int(pid), &wStatus, syscall.WUNTRACED, &usage); err != nil {
//			panic(err)
//		}
//
//		stderrData, err := ioutil.ReadFile("stderr")
//		if err != nil {
//			panic(err)
//		}
//
//		fmt.Printf("||| %#v \n", usage)
//
//		if wStatus.Exited() {
//			//if wStatus.ExitStatus()
//			fmt.Println("fasfaf")
//
//			return Result{
//				ResultCode: RuntimeError,
//				ResultDes:  string(stderrData),
//			}
//		} else {
//			sig := wStatus.Signal()
//			if sig == syscall.SIGXCPU || usage.Utime.Sec > this.TimeLimit || timeLimitKilled {
//				return Result{
//					ResultCode: TimeLimitExceeded,
//					ResultDes:  "",
//				}
//			} else if sig == syscall.SIGXFSZ {
//				return Result{
//					ResultCode: OutputLimitExceeded,
//					ResultDes:  "",
//				}
//			} else if usage.Maxrss > this.MemoryLimit {
//				return Result{
//					ResultCode: MemoryLimitExceeded,
//					ResultDes:  string(stderrData),
//				}
//			}
//		}
//
//	} else {
//		fmt.Println("CHILD")
//
//		var rlimit syscall.Rlimit
//		{
//			rlimit.Cur = uint64(this.TimeLimit)
//			rlimit.Max = uint64(this.TimeLimit)
//			err := syscall.Setrlimit(syscall.RLIMIT_CPU, &rlimit)
//			if err != nil {
//				panic(err)
//			}
//		}
//		{
//			rlimit.Cur = uint64(this.MemoryLimit * 1024)
//			rlimit.Max = uint64(this.MemoryLimit * 1024)
//			err := syscall.Setrlimit(syscall.RLIMIT_AS, &rlimit)
//			if err != nil {
//				panic(err)
//			}
//		}
//		{
//			rlimit.Cur = uint64(this.OutputLimit * 1024)
//			rlimit.Max = uint64(this.OutputLimit * 1024)
//			err := syscall.Setrlimit(syscall.RLIMIT_FSIZE, &rlimit)
//			if err != nil {
//				panic(err)
//			}
//		}
//		{
//			rlimit.Cur = uint64(this.MemoryLimit)
//			rlimit.Max = uint64(this.MemoryLimit)
//			err := syscall.Setrlimit(syscall.RLIMIT_AS, &rlimit)
//			if err != nil {
//				panic(err)
//			}
//		}
//
//		cmd := exec.Command(cmdName, cmdArg...)
//
//		output, err := os.Create(outputFile)
//		if err != nil {
//			panic(err)
//		}
//		cmd.Stdout = bufio.NewWriter(output)
//
//		os.Stdout = output
//		os.Stderr = output
//
//		input, err := os.Open(inFile)
//		if err != nil {
//			panic(err)
//		}
//		cmd.Stdin = bufio.NewReader(input)
//
//		errput, err := os.Create("stderr")
//		if err != nil {
//			panic(err)
//		}
//		cmd.Stderr = bufio.NewWriter(errput)
//
//		err = cmd.Run()
//		if err != nil {
//			panic(err)
//		}
//
//		//var a = 0
//		//for i := 0; i <= 1000544; i++ {
//		//	fmt.Println("afsasf", rand.Float64())
//		//	a++
//		//}
//		//time.Sleep(10 * time.Second)
//		os.Exit(0)
//	}
//
//	return Result{
//		ResultCode: Normal,
//		ResultDes:  "",
//	}
//}
//
//func (this *JudgeBase) Compare() {
//	//syscall.Wait4()
//}
//
//func (this *JudgeBase) Kill(pid int) {
//	syscall.Kill(pid, syscall.SIGKILL)
//}
