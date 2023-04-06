package samples

import (
	"fmt"
	"log"
)

type StringerStruct struct {
	Name string
}

type CustomRaiseError struct {
	ErrorDescription string
}

func (s StringerStruct) String() string {
	var catch string = "human"
	return fmt.Sprintf("print i am %v", catch)
}

func CallStringerStruct() {
	log.Print("call\n")
	s := new(StringerStruct)
	fmt.Printf("value=%v %T\n", s, s)
}

func (e *CustomRaiseError) Error() string {
	return fmt.Sprintf("custom Error! %v", e.ErrorDescription)
}

func CallErrorStruct() error {
	h := false
	if h {
		return nil
	}
	return &CustomRaiseError{ErrorDescription: "Error"}
}

func CallErros() {
	if err := CallErrorStruct(); err != nil {
		log.Fatalln(err)
	}
}
