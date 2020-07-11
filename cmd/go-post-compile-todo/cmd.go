package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/LiamYabou/top100-scrapy/variable"
	"github.com/LiamYabou/top100-pkg/logger"
)

func main() {
	fmt.Printf("app_uri: %s\n", variable.AppURI)
	fmt.Printf("migration_url: %s\n", variable.MigrationURL)
	sourcePath := fmt.Sprintf("file://%s/db/migrations", variable.AppURI)
	m, err := migrate.New(sourcePath, variable.MigrationURL)
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
