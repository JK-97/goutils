package main

import (
	"time"

	"github.com/JK-97/go-utils/logger"
)

func main() {
	logger.Info("Hello")
	logger.Error("Error")

	logWithCallerSkip()

	logWithFields()
}

func hello() {

	logger.Info("Hello")
	logger.Error("Error")
}

func logWithCallerSkip() {

	logger.SetLoggerConfig(logger.Configuration{
		CallerSkip: 1,
	})
	hello()
}

func logWithFields() {
	logger.SetLoggerConfig(logger.Configuration{
		Fields: logger.Fields{"hello": "world", "time": time.Now()},
	})
	hello()
}
