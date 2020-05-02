package app

// Initialize the actions of launching the app,
// and also can load the additional services manually.

import (
  "database/sql"
	_ "github.com/lib/pq"
  "top100-scrapy/pkg/db"
)

var (
  DBconn *sql.DB
  env = os.Getenv("TOP100_ENV")
  file *os.File
)

func init() {
  switch env {
  case "development":
    file, err = logger.SetDevConfigs()
    if err != nil {
      // TODO: Add the error into logger.
    }
  case "staging":
    logger.SetStagingConfigs()
  case "production":
    logger.SetProductionConfigs()
  }

  DBconn, err = db.Open()
  if err != nil {
    // TODO: Add the error into logger.
  }
}
