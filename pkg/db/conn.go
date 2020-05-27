package db

// Initialize the connection of DB.

import (
	"fmt"
	"os"
	"github.com/jackc/pgx/v4/pgxpool"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	dbName     = os.Getenv("TOP100_DB_NAME")
	dbUser     = os.Getenv("TOP100_DB_USER")
	dbPassword = os.Getenv("TOP100_DB_PASSWORD")
	dbPort     = os.Getenv("TOP100_DB_PORT")
	dbHost     = os.Getenv("TOP100_DB_HOST")
	sslMode    = os.Getenv("TOP100_SSL_MODE")
	testDbUrl  = os.Getenv("TOP100_DB_TEST_DSN")
)

func Open() (db *pgxpool.Pool, err error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)
	db, err = pgxpool.Connect(context.Background(), dbURL)
	return db, err
}

func OpenTest() (db *pgxpool.Pool, err error) {
	db, err = pgxpool.Connect(context.Background(), testDbUrl)
	return db, err
}

func OpenPQtest() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", testDbUrl)
	return db, err
} 
