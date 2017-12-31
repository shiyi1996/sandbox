package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"self/judge"
)

func main() {
	os.Chdir("/workspace")

	judgeDataStr := flag.String("judge_data", "", "judge data")
	flag.Parse()

	fmt.Println(*judgeDataStr)

	judger := judge.Judger{}
	json.Unmarshal([]byte(*judgeDataStr), &judger)

	judger.DoJudge()
}
