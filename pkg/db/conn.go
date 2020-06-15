package db

// Initialize the connection of DB.

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
)

func Open() (db *pgxpool.Pool, err error) {
	config, err := pgxpool.ParseConfig(variable.DBURL)
	if err != nil {
		return nil, err
	}
	db, err = pgxpool.ConnectConfig(context.Background(), config)
	return db, err
}

func OpenTest() (db *pgxpool.Pool, err error) {
	db, err = pgxpool.Connect(context.Background(), variable.TestDBURL)
	return db, err
}

func OpenPQtest() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", variable.TestDBURL)
	return db, err
} 
