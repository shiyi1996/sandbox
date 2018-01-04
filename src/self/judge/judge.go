/**
 * Created by shiyi on 2017/11/24.
 * Email: shiyi@fightcoder.com
 */

package judge

type Judge interface {
	Compile() Result
	Run(inputFile string, outputFile string) Result
}

func newJudge(language string, timeLimit int64, memoryLimit int64, outputLimit int64) Judge {
	switch language {
	case "c++":
		return &JudgeCpp{
			JudgeBase{
				TimeLimit:   timeLimit,
				MemoryLimit: memoryLimit,
				OutputLimit: outputLimit,
			},
		}
	case "c":
		return &JudgeC{
			JudgeBase{
				TimeLimit:   timeLimit,
				MemoryLimit: memoryLimit,
				OutputLimit: outputLimit,
			},
		}
	case "python":
		return &JudgePy{
			JudgeBase{
				TimeLimit:   timeLimit,
				MemoryLimit: memoryLimit,
				OutputLimit: outputLimit,
			},
		}
	case "golang":
		return &JudgeGo{
			JudgeBase{
				TimeLimit:   timeLimit,
				MemoryLimit: memoryLimit,
				OutputLimit: outputLimit,
			},
		}
	case "java":
		return &JudgeJava{
			JudgeBase{
				TimeLimit:   timeLimit,
				MemoryLimit: memoryLimit,
				OutputLimit: outputLimit,
			},
		}
	default:
		panic("No such judge: " + language)
	}
}
