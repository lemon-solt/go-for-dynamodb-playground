package samples

import (
	"fmt"
	"time"
)

// 送信と受信
/*
func hoge(first <-chan int, second chan<- int)
*/

func longLongProcess(taskName []int) {
	time.Sleep(2 * time.Second)
	fmt.Printf("finished long process: %v\n", taskName)

}

func taskGen(tasks ...int) [][]int {
	var taskSplit [][]int
	size := len(tasks)

	for i := 0; i < size; i += 3 {
		end := i + 3
		if size < end {
			end = size
		}
		taskSplit = append(taskSplit, tasks[i:end])
	}

	return taskSplit
}

func multiprocess(tasksList [][]int) {

	for _, v := range tasksList {
		go longLongProcess(v)
	}
}

func FanOutMain() {
	start := time.Now()
	done := make(chan interface{})
	defer close(done)

	taskList := []int{1, 2, 3, 4, 5, 6}
	// taskを生成
	tasks := taskGen(taskList...)
	// taskを並列依頼
	multiprocess(tasks)

	fmt.Println("created: ", tasks)

	fmt.Println("finished: ", time.Since(start))
}
