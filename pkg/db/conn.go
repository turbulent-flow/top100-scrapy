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
	dbName     = os.Getenv("DB_NAME")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbPort     = os.Getenv("DB_PORT")
	dbHost     = os.Getenv("DB_HOST")
	sslMode    = os.Getenv("SSL_MODE")
	testDBURL  = os.Getenv("TEST_DB_DSN")
	maxPoolConns = os.Getenv("MAX_POOL_CONNECTIONS")
	minPoolConns = os.Getenv("MIN_POOL_CONNECTIONS")
)

func Open() (db *pgxpool.Pool, err error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=%s&pool_min_conns=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode, maxPoolConns, minPoolConns)
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
	db, err = pgxpool.ConnectConfig(context.Background(), config)
	return db, err
}

func OpenTest() (db *pgxpool.Pool, err error) {
	db, err = pgxpool.Connect(context.Background(), testDBURL)
	return db, err
}

func OpenPQtest() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", testDBURL)
	return db, err
} 
