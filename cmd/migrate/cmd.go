package main

import (
	"os"
	"github.com/LiamYabou/top100-scrapy/variable"
	"github.com/LiamYabou/top100-scrapy/pkg/automation"
	"github.com/LiamYabou/top100-pkg/logger"
)

// # the introduce of the subcommands:
// ## up: migrate up [-step] <number>
// ## down: migrate down [-step] <number>
// ## force: migrate force
// ask help for the `-h` flag, e.g., migrate up -h

func main() {
	err, msg := automation.MigrateDB(variable.MigrationURL, os.Args)
	if err != nil {
		logger.Error(msg, err)
	}
}
