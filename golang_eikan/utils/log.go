package utils

import (
	"io"
	"log"
	"os"
)

// LoggingSetUp is for set up
func LoggingSetUp(logFile string) {
	// RDWRはreadとwrite。パーミッションで0666は読み書きができるユーザーその他。
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

// ErrorLog saves log message
func ErrorLog(message string, err error) {
	log.Printf("%s: %v\n", message, err)
}
