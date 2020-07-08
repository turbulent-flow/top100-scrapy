package main

import (
	"fmt"
	"context"
	"strings"
	"github.com/LiamYabou/top100-pkg/db"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
	"github.com/LiamYabou/top100-pkg/logger"
)

func main() {
	s := fmt.Sprintf("/top100_%s", variable.Env)
	dbURL := strings.ReplaceAll(variable.DBURL, s, "")
	dbPool, err := db.Open(dbURL)
	if err != nil {
		logger.Error("Failed to connect the DB.\n", err)
	}
	defer dbPool.Close()
	stmt := fmt.Sprintf("DROP DATABASE IF EXISTS top100_%s", variable.Env)
	_, err = dbPool.Exec(context.Background(), stmt)
	if err != nil {
		logger.Error("Failed to drop a database.", err)
	}
	stmt = fmt.Sprintf("CREATE DATABASE top100_%s", variable.Env)
	_, err = dbPool.Exec(context.Background(), stmt)
	if err != nil {
		logger.Error("Failed to create a database.", err)
	}
	// start a new connection that indicates the database, so that we can create the dedicated extension for the current database.
	secondDBpool, err := db.Open(variable.DBURL)
	if err != nil {
		logger.Error("Failed to connect the DB.\n", err)
	}
	stmt = "CREATE EXTENSION IF NOT EXISTS ltree"
	_, err = secondDBpool.Exec(context.Background(), stmt)
	if err != nil {
		logger.Error("Failed to create a extension.", err)
	}
	defer secondDBpool.Close()
	fmt.Println("  > Done.")
}
