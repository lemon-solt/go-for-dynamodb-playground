package samples

import "fmt"

func HelloInterface(itf interface{}) {
	switch v := itf.(type) {
	case int:
		fmt.Print("int", v)
	default:
		fmt.Print("not int", v)
	}
}
