package samples

import (
	"fmt"
)

// func WorkersPettern() {
// 	var Ball int
// 	table := make(chan int)
// 	go player(table, "A")
// 	go player(table, "B")

// 	table <- Ball
// 	time.Sleep(10 * time.Second)
// 	<-table
// }

// func player(table chan int, id string) {
// 	for {
// 		fmt.Println("ball get: ", id)
// 		ball := <-table
// 		ball++
// 		time.Sleep(100 * time.Millisecond)
// 		table <- ball
// 	}
// }

// func producerx(ch chan int, d time.Duration) {
// 	var i int
// 	for {
// 		ch <- i
// 		i++
// 		time.Sleep(d)
// 	}
// }

// func reader(out chan int) {
// 	for x := range out {
// 		fmt.Println(x)
// 	}
// }

// func WorkersPettern() {
// 	ch := make(chan int)
// 	out := make(chan int)
// 	go producerx(ch, 100*time.Millisecond)
// 	go producerx(ch, 250*time.Millisecond)
// 	go reader(out)
// 	for i := range ch {
// 		out <- i
// 	}
// }

// func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for {
// 		task, ok := <-tasksCh
// 		if !ok {
// 			return
// 		}
// 		d := time.Duration(task) * time.Millisecond
// 		time.Sleep(d)
// 		fmt.Println("processing task", task)
// 	}
// }

// func pool(wg *sync.WaitGroup, workers, tasks int) {
// 	tasksCh := make(chan int)

// 	for i := 0; i < workers; i++ {
// 		go worker(tasksCh, wg)
// 	}

// 	for i := 0; i < tasks; i++ {
// 		tasksCh <- i
// 	}

// 	close(tasksCh)
// }

// func WorkersPettern() {
// 	var wg sync.WaitGroup
// 	wg.Add(36)
// 	go pool(&wg, 36, 50)
// 	wg.Wait()
// }

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func WorkersPettern() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
