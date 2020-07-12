package main

import (
	"os"
	"github.com/LiamYabou/top100-scrapy/variable"
	"github.com/LiamYabou/top100-scrapy/pkg/automation"
	"github.com/LiamYabou/top100-pkg/logger"
)

func main() {
	err, msg := automation.MigrateDB(variable.TestDBURL, os.Args)
	if err != nil {
		logger.Error(msg, err)
	}
}
