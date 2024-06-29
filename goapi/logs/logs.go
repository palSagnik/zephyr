package logs

import (
	"io"
	"log"
	"os"
)

func CustomLogger() {
	errorFile, _ := os.OpenFile("./logs/error.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	mw := io.MultiWriter(os.Stdout, errorFile)
	
	log.SetOutput(mw)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}