package main

import (
	"top100-scrapy/pkg/app"
	"top100-scrapy/pkg/logger"
)

func main() {
	logger.Debug("Debug starts - enqueue insert categories job.")
	defer app.Finalize()
	performJob()
	logger.Debug("Debug stops -  enqueue insert categories job.")
}

func performJob() {
	ch, err := app.AMQPconn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel.", err)
	}
	defer ch.Close()
}
