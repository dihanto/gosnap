package config

import (
	"log"
	"os"
)

func InitLogFile() *os.File {
	logFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	return logFile
}
