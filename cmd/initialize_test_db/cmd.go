package main

import (
	"github.com/LiamYabou/top100-scrapy/pkg/automation"
	"github.com/LiamYabou/top100-pkg/logger"
)

func main() {
	err := automation.InitTestDB()
	if err != nil {
		logger.Error("Could not initialize the DB.", err)
	}
}
