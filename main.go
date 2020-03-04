package main

import (
	"gitlab.jiangxingai.com/applications/base-modules/internal-sdk/go-utils/logger"
)

func main() {
	logger.Info("Hello")
	logger.Error("Error")

	logger.SetLoggerConfig(logger.Configuration{
		CallerSkip: 1,
	})
	hello()
}

func hello() {

	logger.Info("Hello")
	logger.Error("Error")
}
