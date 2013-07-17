package log

import (
	"fmt"
	"log"
	"os"
)

var logFile, logFileErr = os.OpenFile("/var/www/go/logs/go_thy.log", os.O_RDWR|os.O_APPEND, 0660)

//var Log = log.New(LogFile, "", log.Ldate|log.Ltime)
var infoLog = log.New(logFile, "Info: ", log.Ldate|log.Ltime)
var debugLog = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime)

func InitLog() {
	if logFileErr != nil {
		fmt.Println("  Error initializing shared log: ", logFileErr)
	} else {
		fmt.Println("  Logging to go_thy.log ")
	}
}

func Debug(err error, message string) {
	if err != nil {
		debugLog.Println("DEBUG : ", message, err)
	}
}

func Info(a ...interface{}) {
	infoLog.Println(a...)
}
