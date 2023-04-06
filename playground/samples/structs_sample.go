package samples

import (
	"fmt"
	"strconv"
)

type SampleStruct struct {
	Name     string
	Age      int
	Favarite string
}

type AddSampleStruct struct {
	SampleStruct
	Other string
}

func (s SampleStruct) Iam() {
	fmt.Printf("Hi. I'm %v", s.Name)
}

func (a AddSampleStruct) MyProfile() {
	fmt.Printf("I'm %v %v", a.Name, a.Other)
}

func (a AddSampleStruct) Concat() {
	customAge := strconv.Itoa(a.Age) + " hi!"
	fmt.Printf("I'm %v", customAge)
}
