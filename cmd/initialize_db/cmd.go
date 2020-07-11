package main

import (
	"github.com/LiamYabou/top100-scrapy/variable"
	"github.com/LiamYabou/top100-scrapy/automation"
	"github.com/LiamYabou/top100-pkg/logger"
)

func main() {
	err := automation.InitDB(variable.Env)
	if err != nil {
		logger.Error("Could not initialize the DB.", err)
	}
	automation.Finalize()
}
