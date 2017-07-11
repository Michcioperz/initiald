package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var logFile *os.File

func initLogging() {
	defer log.Println("logging initialized")
	initialTime := time.Now().Unix()
	fileName := fmt.Sprintf("initiald.%v.log", initialTime)
	filePath1 := path.Join("/var", "log", fileName)
	var o1err error
	logFile, o1err = os.OpenFile(filePath1, os.O_APPEND|os.O_CREATE, 0644)
	if o1err != nil {
		defer log.Println("log path ", filePath1, " was not used due to error ", o1err)
		filePath2 := path.Join("/", fileName)
		var o2err error
		logFile, o2err = os.OpenFile(filePath2, os.O_APPEND|os.O_CREATE, 0644)
		if o2err != nil {
			defer log.Println("log path ", filePath2, " was not used due to error ", o2err)
			defer log.Println("log file not open, falling back to stderr alone")
			logFile = nil
			log.SetOutput(os.Stderr)
			return
		}
	}
	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
}

func closeLogging() {
	log.Println("terminating logging")
	if logFile != nil {
		logFile.Close()
	}
}
