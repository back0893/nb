package utils

import "log"

var LoggerObject *Logger

func init() {
	LoggerObject = &Logger{}
}

type Logger struct{}

func (Logger) Write(errStr string) {
	log.Println("错误=>", errStr)
}
