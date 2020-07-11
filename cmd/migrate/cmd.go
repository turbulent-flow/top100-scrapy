package main

import (
	"os"
	"errors"
	"fmt"
	"flag"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/LiamYabou/top100-scrapy/v2/variable"
	"github.com/LiamYabou/top100-pkg/logger"
)

// # the introduce of the subcommands:
// ## up: migrate up [-step] <number>
// ## down: migrate down [-step] <number>
// ## force: migrate force
// ask help for the `-h` flag, e.g., migrate up -h

func main() {
	if len(os.Args) < 2 {
		content := "Expected the \"up\" or \"down\" or \"force\" subcommands."
		logger.Error("An error occured.", errors.New(content))
	}
	upCmd := flag.NewFlagSet("up", flag.ExitOnError)
	upStep := upCmd.Int("step", 0, "If the number is zero, the DB will be migrated to the last version according to the resources of the \"db\\migrations\" directory. You can speicfy the version of the migration which you expect to migrate to.")
	downCmd := flag.NewFlagSet("down", flag.ExitOnError)
	downStep := downCmd.Int("step", 0, "If the number is zero, the DB will be migrated all the way down according the the resources of the \"db\\migrations\" directory. You can speicfy the version of the migration which you expect to migrate to.")
	forceCmd := flag.NewFlagSet("force", flag.ExitOnError)
	forceCmd.Bool("enable", true, "The state of the \"dirty\" column of the \"schema_migrations\" table is set to false. It works by cleaning the dirty state first when you met the error: \"Dirty database version x. Fix and force version...\" during the migration, and then migrate all the way down.")
	sourcePath := fmt.Sprintf("file://%s/db/migrations", variable.AppURI)
	m, err := migrate.New(sourcePath, variable.MigrationURL)
	if err != nil {
		logger.Error("Failed to establish the connection of the migration.", err)
	}
	defer m.Close()
	switch os.Args[1] {
	case "up":
		upCmd.Parse(os.Args[2:])
		if *upStep == 0 {
			err = m.Up()
		} else if *upStep > 0 {
			err = m.Steps(*upStep)
		} else {
			content := "Expected the number is not less than zero."
			logger.Error("An error occured.", errors.New(content))
		}
	case "down":
		downCmd.Parse(os.Args[2:])
		if *downStep == 0 {
			err = m.Down()
		} else if *downStep > 0 {
			n := *downStep * -1
			err = m.Steps(n)
		} else {
			content := "Expected the number is not less than zero."
			logger.Error("An error occured.", errors.New(content))
		}
	case "force":
		forceCmd.Parse(os.Args[2:])
		version, _, err := m.Version()
		if err != nil && err.Error() == "no migration" {
			fmt.Println("  > NOTE: There is no migration.")
		} else if err != nil {
			logger.Error("Failed to fetch the version of the migration.", err)
		}
		err = m.Force(int(version))
	default:
		content := "Expected the \"up\" or \"down\" or \"force\" subcommands."
		logger.Error("An error occured.", errors.New(content))
	}
	if err != nil && err.Error() == "no change" {
		fmt.Println("  > NOTE: There is no change related to the operation of the migration.")
		return
	} else if err != nil {
		logger.Error("Failed to establish the connection of the migration.", err)
	}
	fmt.Println("  > Done.")
}
