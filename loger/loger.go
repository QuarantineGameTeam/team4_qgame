package loger

import (
	"log"
	"os"
	"time"
)

var (
	outfile, _ = os.OpenFile("logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile    = log.New(outfile, "", 0)
)

func ForrError(err error, v ...interface{}) {
	if err != nil {
		LogFile.Fatalln("Time :", time.Now(), v, err)
	}
}
