package logger

// Set the configurations of logger, which vary with the different environment.

import (
  "os"
  "fmt"
  log "github.com/sirupsen/logrus"
)

var (
  appUri = os.Getenv("APP_URI")
)

func SetDevConfigs() (file *os.File, err error) {
  filePath := fmt.Sprintf("%s/logs/development.log", appUri)
  file, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
  log.SetOutput(file)
  log.SetFormatter(&log.JSONFormatter{})
  log.SetLevel(log.DebugLevel)
  return file, err
}

func SetStagingConfigs() {}

func SetProductionConfigs() {}
