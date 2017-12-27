package main

import (
	"flag"

	"self/judge"
)

func main() {
	judgeType := flag.String("judge_type", "default", "judge type: dafault, contest, test, special")
	language := flag.String("language", "cpp", "the language of the code file")
	codeFile := flag.String("code_file", "code.cpp", "code file")
	timeLimit := flag.Int64("time_limit", 0, "time limit")
	memoryLimit := flag.Int64("memory_limit", 0, "memory limit")
	outputLimit := flag.Int64("output_limit", 0, "output limit")

	flag.Parse()

	judge.DoJudge(*judgeType, *language, *codeFile, *timeLimit, *memoryLimit, *outputLimit)
}
