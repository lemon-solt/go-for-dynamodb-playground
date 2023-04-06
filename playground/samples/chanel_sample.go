package samples

import "fmt"

func loopProc(items []int, c chan int) {
	s := 0
	for _, v := range items {
		s += v
	}
	c <- s
}

func CallChanel() {
	items := []int{1, 2, 3, 4, 5}
	chanel := make(chan int)
	go loopProc(items, chanel)
	x := <-chanel
	fmt.Println(x)
}

func BufferdChanel() {
	ch := make(chan int, 2)
	ch <- 200
	ch <- 100
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}
}

func loopProcCurrent(items []int, c chan int) {
	s := 0
	for _, v := range items {
		s += v
		c <- s
	}
	close(c)
}

func CallGetForChanel() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go loopProcCurrent(s, c)

	for i := range c {
		fmt.Println(i)
	}

}
