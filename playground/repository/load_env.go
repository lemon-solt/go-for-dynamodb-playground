package repository

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type SettingList struct {
	Region          string
	Aws_seacret_key string
	Aws_access_key  string
}

var EnvSetting = new(SettingList)

func LoadEnv() {
	err := godotenv.Load("../.go_env")
	if err != nil {
		fmt.Printf("env read failed: %v", err)
	}

	EnvSetting.Aws_access_key = os.Getenv("aws_access_key_id")
	EnvSetting.Region = os.Getenv("region")
	EnvSetting.Aws_seacret_key = os.Getenv("aws_secret_access_key")

}
