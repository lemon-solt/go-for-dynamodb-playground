package samples

import (
	"fmt"
	"time"
)

func CallDefaultForBreak() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)

OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("hi")
		case <-boom:
			fmt.Println("end!")
			break OuterLoop
		default:
			fmt.Println("....")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
