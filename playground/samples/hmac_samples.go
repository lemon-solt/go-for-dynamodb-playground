package samples

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Server(apiKey, sign string, data []byte) {
	var DB = map[string]string{
		"hoge": "seacretHoge",
	}
	apiSecret := DB[apiKey]

	h := hmac.New(sha256.New, []byte(apiSecret))

	h.Write(data)
	expect := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign == expect)
}

func CallHmacSample() {
	apiKey := "hoge"
	seacretKey := "seacretHoge"

	data := []byte(`{"key": 1}`)

	h := hmac.New(sha256.New, []byte(seacretKey))
	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign)

	Server(apiKey, sign, data)
}
