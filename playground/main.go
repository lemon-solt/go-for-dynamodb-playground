package main

import (
	"fmt"
	"playground/repository"
	"playground/samples"
	"playground/utils"
)

func init() {
	repository.LoadEnv()
	utils.LoggingSettings(utils.Config.LogFile)
}

func main() {
	fmt.Println("envsetting: ", repository.EnvSetting)
	samples.CallHmacSample()

}
