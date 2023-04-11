package samples

import (
	"fmt"
	"math/rand"
	"time"
)

func tProc(tasks []int) <-chan int {
	tChan := make(chan int)

	go func() {
		defer close(tChan)
		var value int
		for _, v := range tasks {
			value += v
		}
		time.Sleep(waitTime * time.Second)
		tChan <- value
	}()

	return tChan
}

func tFanin(chanels []<-chan int) <-chan int {
	sumChan := make(chan int)

	fnin := func(c <-chan int) {
		for {
			v, ok := <-c
			if !ok {
				return
			}

			sumChan <- v
		}
	}

	for _, c := range chanels {
		go fnin(c)
	}

	return sumChan

}

func allProc(stream <-chan int, count int) {

	for i := 0; i < count; i++ {
		fmt.Println(<-stream)
	}
}

func generateStacks() (stacks [][]int) {

	t := func(num int) []int {
		var s []int
		for i := 0; i < 3; i++ {
			rn := rand.Intn(100)
			s = append(s, i*rn)
		}
		return s
	}

	for i := 0; i < taskArrays; i++ {
		stacks = append(stacks, t(overCaluc))
	}

	return stacks
}

func startFanOutA() {
	// two task

	multitask := generateStacks() // [[1,2,3], [5,6,7],,,,,,]
	finders := make([]<-chan int, len(multitask))
	for i, v := range multitask {
		finders[i] = tProc(v)
	}

	faninStream := tFanin(finders)

	allProc(faninStream, len(multitask))
}

const (
	waitTime   = 2
	taskArrays = 10
	overCaluc  = 200
)

func FanOutSampleA() {
	fmt.Println("start")
	start := time.Now()
	startFanOutA()

	fmt.Println(time.Since(start))
}
