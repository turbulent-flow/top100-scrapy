package automation

import (
	"fmt"
	"context"
	"strings"
	"github.com/LiamYabou/top100-pkg/db"
	"github.com/LiamYabou/top100-scrapy/variable"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DBpool *pgxpool.Pool
	TestDBpool *pgxpool.Pool
)

func InitDB(env string) (err error) {
	s := fmt.Sprintf("/top100_%s", env)
	dbURL := strings.ReplaceAll(variable.DBURL, s, "")
	DBpool, err = db.Open(dbURL)
	defer DBpool.Close()
	stmt := fmt.Sprintf("DROP DATABASE IF EXISTS top100_%s", env)
	_, err = DBpool.Exec(context.Background(), stmt)
	stmt = fmt.Sprintf("CREATE DATABASE top100_%s", env)
	_, err = DBpool.Exec(context.Background(), stmt)
	return
}

func InitTestDB() (err error) {
	s := "/top100_test"
	dbURL := strings.ReplaceAll(variable.TestDBURL, s, "")
	TestDBpool, err = db.Open(dbURL)
	defer TestDBpool.Close()
	stmt := "DROP DATABASE IF EXISTS top100_test"
	_, err = TestDBpool.Exec(context.Background(), stmt)
	stmt = "CREATE DATABASE top100_test"
	_, err = TestDBpool.Exec(context.Background(), stmt)
	return
}
