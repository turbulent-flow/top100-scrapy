package db

// Initialize the connection of DB.

import (
  "fmt"
  "os"
  "database/sql"
	_ "github.com/lib/pq"
)

var (
  dbName = os.Getenv("TOP100_DB_NAME")
  dbUser = os.Getenv("TOP100_DB_USER")
  dbPassword = os.Getenv("TOP100_DB_PASSWORD")
  dbPort = os.Getenv("TOP100_DB_PORT")
  dbHost = os.Getenv("TOP100_DB_HOST")
  sslMode = os.Getenv("TOP100_SSL_MODE")
)

func Open() (db *sql.DB, err error) {
  dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)
  db, err = sql.Open("postgres", dbUrl)
  return db, err
}
