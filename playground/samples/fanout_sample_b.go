package samples

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func getTask(done <-chan interface{}, taskId string) <-chan string {
	defer fmt.Println("### getTask has finished")
	currentTaskStream := make(chan string)

	go func() {
		defer close(currentTaskStream)
		for {
			select {
			case <-done:
				return
			case currentTaskStream <- taskId:
			}
		}
	}()

	return currentTaskStream
}

func fanInFunc(
	done <-chan interface{},
	channels ...<-chan string,
) <-chan string {

	var wg sync.WaitGroup
	faninStream := make(chan string)

	taskStart := func(c <-chan string) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case faninStream <- i + "_after_proc":
				time.Sleep(3 * time.Second)
			}
		}
	}
	wg.Add(len(channels))
	fmt.Println("counter: ", len(channels))
	for _, c := range channels {
		go taskStart(c)
	}

	go func() {
		wg.Wait()
		close(faninStream)
	}()
	return faninStream
}

func sumOut(
	done <-chan interface{},
	valueStream <-chan string,
) <-chan string {
	defer fmt.Println("=== SumOut finished")
	changeStream := make(chan string)
	fmt.Println("=== SumOut has started")
	go func() {
		defer close(changeStream)
		for {
			select {
			case <-done:
				return
			case ok := <-valueStream:
				changeStream <- ok
			}
		}
	}()
	return changeStream
}

func startProcessing() {
	fmt.Println("*** startSample has started")
	done := make(chan interface{})
	defer func() {
		fmt.Println("*** startSample has finished")
		close(done)
	}()
	// numFinders := runtime.NumCPU()
	numFinders := 2
	finders := make([]<-chan string, numFinders)

	for i := 0; i < numFinders; i++ {
		finders[i] = getTask(done, strconv.Itoa(i)+"_id")
	}

	for out := range sumOut(done, fanInFunc(done, finders...)) {
		fmt.Println("out...", out)
	}
	// キューの準備とキューの作成
	// キューを各種タスクに振り分け
	// 結果を集約
	// 処理を出力
}

func LetsFanoutFanin() {
	start := time.Now()
	// スタート
	startProcessing()
	fmt.Printf("total execution time: %v\n", time.Since(start))
}
