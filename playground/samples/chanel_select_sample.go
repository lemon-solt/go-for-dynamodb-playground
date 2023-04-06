package samples

import (
	"fmt"
	"time"
)

func gorouting1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second)
	}
}

func gorouting2(ch chan int) {
	for {
		ch <- 1
		time.Sleep(1 * time.Second)
	}
}

func CallSelectChan() {
	c1 := make(chan string)
	c2 := make(chan int)
	go gorouting1(c1)
	go gorouting2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
