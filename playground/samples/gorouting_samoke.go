package samples

import (
	"fmt"
	"sync"
)

func gorouting(s string, wg *sync.WaitGroup) {
	hellos := []map[string]string{}

	hellos = append(hellos, map[string]string{"key": "hello"})
	hellos = append(hellos, map[string]string{"key": "hello2"})
	hellos = append(hellos, map[string]string{"key": "hello3"})
	hellos = append(hellos, map[string]string{"key": "hello4"})

	for i := range hellos {
		fmt.Println("あいさつ", hellos[i]["key"])
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("groutin: ", s)
	// }
	wg.Done()
}

func CallGoroutin() {
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go gorouting("hello", &wg)

	for i := 0; i < 10; i++ {
		fmt.Println("out loop")
	}

}
