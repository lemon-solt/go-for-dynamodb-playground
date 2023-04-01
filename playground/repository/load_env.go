package repository

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type SettingList struct {
	Region        string
	AwsSeacretKey string
	AwsAccessKey  string
}

var EnvSetting = new(SettingList)

func LoadEnv() {
	err := godotenv.Load("../.go_env")
	if err != nil {
		fmt.Printf("env read failed: %v", err)
	}

	EnvSetting.AwsAccessKey = os.Getenv("aws_access_key_id")
	EnvSetting.Region = os.Getenv("region")
	EnvSetting.AwsSeacretKey = os.Getenv("aws_secret_access_key")

}
