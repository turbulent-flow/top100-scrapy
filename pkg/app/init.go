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
)

func init() {
  DBconn, err = db.Open()
  if err != nil {
    // TODO: Add the error into logger.
  }
}
