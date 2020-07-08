package automation

import (
	"fmt"
	"context"
	"strings"
	"github.com/LiamYabou/top100-pkg/db"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DBpool *pgxpool.Pool
	SecondDBpool *pgxpool.Pool
)

func InitDB(env string) (err error) {
	s := fmt.Sprintf("/top100_%s", env)
	dbURL := strings.ReplaceAll(variable.DBURL, s, "")
	DBpool, err = db.Open(dbURL)
	stmt := fmt.Sprintf("DROP DATABASE IF EXISTS top100_%s", env)
	_, err = DBpool.Exec(context.Background(), stmt)
	stmt = fmt.Sprintf("CREATE DATABASE top100_%s", env)
	_, err = DBpool.Exec(context.Background(), stmt)
	// start a new connection that indicates the database, so that we can create the dedicated extension for the current database.
	SecondDBpool, err = db.Open(variable.DBURL)
	stmt = "CREATE EXTENSION IF NOT EXISTS ltree"
	_, err = SecondDBpool.Exec(context.Background(), stmt)
	return
}