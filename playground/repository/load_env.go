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

	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load("../.go_env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	EnvSetting.Aws_access_key = os.Getenv("aws_access_key_id")
	EnvSetting.Region = os.Getenv("region")
	EnvSetting.Aws_seacret_key = os.Getenv("aws_secret_access_key")

}
