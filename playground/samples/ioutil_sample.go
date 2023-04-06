package samples

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func ReadIoutil() {
	content, err := ioutil.ReadFile("./samples/mutex_map.go")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(content))
}

func ReadIoutilBuffer() {
	read := bytes.NewBuffer([]byte("abc"))
	content, _ := ioutil.ReadAll(read)
	fmt.Println(string(content))
}

func WriteIoutil() {
	ioutil.WriteFile("hoge.html", []byte("<h1>愛はあるんか</h1>"), 0666)
}
