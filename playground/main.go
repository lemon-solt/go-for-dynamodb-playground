package main

import "playground/samples"

func init() {
	// repository.LoadEnv()
	// utils.LoggingSettings(utils.Config.LogFile)
}

func main() {
	// fmt.Println("envsetting: ", repository.EnvSetting)
	// samples.FanOutFanInSample()
	samples.FanOutSampleA()
	// fmt.Println(runtime.NumCPU())

}
