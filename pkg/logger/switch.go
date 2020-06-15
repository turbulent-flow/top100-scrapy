package logger

// Switch the level according with the different environment.

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var env = os.Getenv("TOP100_ENV")

func switchError(entry *log.Entry, msg string) {
	switch env {
	case "development":
		entry.Panic(msg)
	case "staging":
		entry.Panic(msg)
	case "production":
		entry.Fatal(msg)
	}
}
