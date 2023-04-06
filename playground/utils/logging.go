package utils

import (
	"log"
	"os"
)

type ConfigList struct {
	LogFile string
}

var Config ConfigList

func LoggingSettings(logFile string) {
	// logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.SetOutput(multiLogFile)

	// 引数出力
	cmd := os.Args
	for _, v := range cmd {
		log.Println("cmd引数", v)
	}

}
