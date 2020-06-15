package logger

// Switch the level according with the different environment.

import (
	log "github.com/sirupsen/logrus"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
)


func switchError(entry *log.Entry, msg string) {
	switch variable.Env {
	case "development":
		entry.Panic(msg)
	case "staging":
		entry.Panic(msg)
	case "production":
		entry.Fatal(msg)
	}
}
