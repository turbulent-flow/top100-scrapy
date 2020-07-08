package main

import (
	"github.com/LiamYabou/top100-scrapy/v2/variable"
	"github.com/LiamYabou/top100-scrapy/v2/automation"
	"github.com/LiamYabou/top100-pkg/logger"
)

func main() {
	err := automation.InitDB(variable.Env)
	if err != nil {
		logger.Error("Could not initialize the DB.", err)
	}
	automation.Finalize()
}
