package samples

import (
	"fmt"
	"time"
)

func lisGen(i int) (countners []int) {
	for c := 0; c < i; c++ {
		countners = append(countners, c)
	}
	return
}

func CallLoopGoroutin() {

	target := lisGen(10)
	// s := []int{}

	ch := make(chan int, len(target))

	// var wg sync.WaitGroup
	// wg.Add(len(target))
	for _, v := range target {
		go func(v int, c chan<- int) {
			// defer wg.Done()
			// s = append(s, v)
			c <- v
		}(v, ch)
	}

	// wg.Wait()

	// close(ch)
	for {
		select {
		case vv := <-ch:
			fmt.Println(vv, "を受信")
		default:
			time.Sleep(1 * time.Second)
			// fmt.Println("nothing...")
		}
	}
}

func CallBufferChanel() {
	done := make(chan interface{})

	gen := func(done <-chan interface{}) <-chan int {
		res := make(chan int)

		go func() {
			defer close(res)

		l:
			for {
				select {
				case <-done:
					break l
				case res <- 1:
				}
			}
		}()
		return res
	}

	result := gen(done)
	for i := 0; i < 5; i++ {
		fmt.Println(result, "!")
	}

	close(done)
}
