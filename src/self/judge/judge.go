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
	var jd Judge
	switch language {
	case "c++":
		jd = &JudgeCpp{
			JudgeBase{
				TimeLimit:   timeLimit,
				MemoryLimit: memoryLimit,
				OutputLimit: outputLimit,
			},
		}
	case "c":
		jd = &JudgeC{
			JudgeBase{
				TimeLimit:   timeLimit,
				MemoryLimit: memoryLimit,
				OutputLimit: outputLimit,
			},
		}
	case "python":
		jd = &JudgePy{
			JudgeBase{
				TimeLimit:   timeLimit,
				MemoryLimit: memoryLimit,
				OutputLimit: outputLimit,
			},
		}
	default:
		panic("No such judge: " + language)
	}

	return jd
}
