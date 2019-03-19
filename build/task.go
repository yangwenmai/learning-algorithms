package build

import (
	"log"

	"github.com/yangwenmai/learning-algorithms/util"
)

const (
	taskFile = "tasks.txt"
)

type TaskBuilder struct {
	*Builder
	Prefix      string
	First, Last int
}

func (t *TaskBuilder) Build() {
	log.Println("开始，生成新任务")

	res := ""
	count := 0
	ps := t.LC.Problems
	var isWanted func(int) bool
	collect := func() {
		for i := t.First; i <= t.Last && i < len(ps); i++ {
			if !isWanted(i) {
				log.Println("暂时不想做...")
				continue
			}
			res += ps[i].DidaTask(t.Prefix) + "\n"
			count++
		}
	}

	// 根据 prefix 的不同，设置不同的 isWanted
	switch t.Prefix {
	case "do": // to do
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && !ps[i].IsAccepted && !ps[i].IsPaid && !ps[i].HasNoGoOption
		}
	case "re": // review
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && ps[i].IsAccepted && !ps[i].IsPaid
		}
	case "mi": // miss
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && ps[i].IsAccepted && ps[i].HasNoGoOption
		}
	case "fa": // favor
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && ps[i].IsFavor
		}
	default:
		// 和 do 的一样
		isWanted = func(i int) bool {
			return ps[i].ID > 0 && !ps[i].IsAccepted && !ps[i].IsPaid && !ps[i].HasNoGoOption
		}
	}

	collect()
	log.Printf("完成，一共生成了 %d 条新任务\n", count)
	util.WriteFile(taskFile, res)
}
