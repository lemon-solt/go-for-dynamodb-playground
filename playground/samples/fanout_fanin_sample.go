package samples

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func getTimeUTC(
	done <-chan interface{},
	fn func() time.Time,
) <-chan time.Time {
	defer fmt.Println("### getUtcTime has finished")
	valueStream := make(chan time.Time)
	fmt.Println("### getUtcTime has started")
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

func convertTimeJST(
	done <-chan interface{},
	valueStream <-chan time.Time,
	loopNum int,
) <-chan time.Time {
	defer fmt.Println("=== convertJstTime has finished")
	changeStream := make(chan time.Time)
	fmt.Println("=== convertJstTime has started")
	go func() {
		defer close(changeStream)
		for count := 0; count < loopNum; count++ {
			select {
			case <-done:
				return
			case nowUTC := <-valueStream:
				jst := time.FixedZone("Asia/Tokyo", 9*60*60)
				nowJST := nowUTC.In(jst)
				changeStream <- nowJST
			}
		}
	}()
	return changeStream
}

func nowTime() time.Time {
	data := time.Now().UTC()
	time.Sleep(1 * time.Second)
	return data
}

func fanIn(
	done <-chan interface{},
	channels ...<-chan time.Time,
) <-chan time.Time {
	var wg sync.WaitGroup
	rendezvousStream := make(chan time.Time)

	rendezvous := func(c <-chan time.Time) {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		for i := range c {
			select {
			case <-done:
				return
			case rendezvousStream <- i:
			}
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go rendezvous(c)
	}

	go func() {
		wg.Wait()
		close(rendezvousStream)
	}()
	return rendezvousStream
}

func startSample() {
	fmt.Println("*** startSample has started")
	done := make(chan interface{})
	defer func() {
		fmt.Println("*** startSample has finished")
		close(done)
	}()
	numFinders := runtime.NumCPU()
	finders := make([]<-chan time.Time, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = getTimeUTC(done, nowTime)
	}

	for nowJST := range convertTimeJST(done, fanIn(done, finders...), 5) {
		fmt.Println(nowJST)
	}
}

func FanOutFanInSample() {
	start := time.Now()
	startSample()
	fmt.Printf("total execution time: %v\n", time.Since(start))
}
