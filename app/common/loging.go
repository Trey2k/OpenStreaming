package common

import (
	"io"
	"log"
	"os"
)

type LogStruct struct {
	Info  *log.Logger
	Error *log.Logger
}

var Loggers LogStruct

func init() {
	infoFile, err := os.OpenFile("/root/logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	errorFile, err := os.OpenFile("/root/logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	Loggers.Info = log.New(io.MultiWriter(os.Stdout, infoFile), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Loggers.Error = log.New(io.MultiWriter(os.Stderr, errorFile), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
