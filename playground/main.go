package main

import (
	"fmt"
	"playground/repository"
)

func init() {
	repository.LoadEnv()
}

func main() {
	fmt.Println("envsetting: ", repository.EnvSetting)
}
