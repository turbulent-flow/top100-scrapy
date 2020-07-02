package app

import (
  "time"
  "github.com/LiamYabou/top100-scrapy/v2/variable"
  "github.com/getsentry/sentry-go"
)

func Finalize() {
  if variable.Env == "development" {
    logFile.Close()
  }
  DBpool.Close()
  AMQPconn.Close()
  sentry.Recover() // Capture the unhandled panic
  sentry.Flush(5 * time.Second) // Set the timeout to the maximum duration the program can afford to wait.
}
