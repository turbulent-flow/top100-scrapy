package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
	"github.com/LiamYabou/top100-pkg/logger"
)

func main() {
	sourcePath := fmt.Sprintf("file://%s/db/migrations", variable.AppURI)
	m, err := migrate.New(sourcePath, variable.TestDBURL)
	if err != nil {
		logger.Error("Failed to establish the connection of the migration.", err)
	}
	err = m.Up()
	if err != nil && err.Error() == "no change" {
		fmt.Println("  > NOTE: There is no change related to the operation of the migration.")
		return
	} else if err != nil {
		logger.Error("Failed to establish the connection of the migration.", err)
	}
	fmt.Println("  > Done.")
}