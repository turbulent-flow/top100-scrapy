package automation

import (
	"fmt"
	"errors"
	"flag"
	"github.com/LiamYabou/top100-scrapy/variable"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// # the introduce of the subcommands:
// ## up: migrate up [-step] <number>
// ## down: migrate down [-step] <number>
// ## force: migrate force
// ask help for the `-h` flag, e.g., migrate up -h
func MigrateDB(migrationURL string, args []string) (err error, message string) {
	if len(args) < 2 {
		content := "Expected the \"up\" or \"down\" or \"force\" subcommands."
		return errors.New(content), "An eror occured."
	}
	upCmd := flag.NewFlagSet("up", flag.ExitOnError)
	upStep := upCmd.Int("step", 0, "If the number is zero or the flag is not set, the DB will be migrated to the last version according to the resources of the \"db\\migrations\" directory. You can speicfy the version of the migration which you expect to migrate to.")
	downCmd := flag.NewFlagSet("down", flag.ExitOnError)
	downStep := downCmd.Int("step", 0, "If the number is zero or the flag is not set, the DB will be migrated all the way down according the the resources of the \"db\\migrations\" directory. You can speicfy the version of the migration which you expect to migrate to.")
	forceCmd := flag.NewFlagSet("force", flag.ExitOnError)
	forceCmd.Bool("enable", true, "The default value of this flag is true. The state of the \"dirty\" column of the \"schema_migrations\" table is set to false. It works by cleaning the dirty state first when you met the error: \"Dirty database version x. Fix and force version...\" during the migration, and then you should migrate all the way down manually.")
	sourcePath := fmt.Sprintf("file://%s/db/migrations", variable.AppURI)
	m, err := migrate.New(sourcePath, migrationURL)
	if err != nil {
		return err, "Failed to establish the connection of the migration."
	}
	defer m.Close()
	switch args[1] {
	case "up":
		upCmd.Parse(args[2:])
		if *upStep == 0 {
			err = m.Up()
		} else if *upStep > 0 {
			err = m.Steps(*upStep)
		} else {
			content := "Expected the number is not less than zero."
			return errors.New(content), "An eror occured."
		}
	case "down":
		downCmd.Parse(args[2:])
		if *downStep == 0 {
			err = m.Down()
		} else if *downStep > 0 {
			n := *downStep * -1
			err = m.Steps(n)
		} else {
			content := "Expected the number is not less than zero."
			return errors.New(content), "An eror occured."
		}
	case "force":
		forceCmd.Parse(args[2:])
		version, _, err := m.Version()
		if err != nil && err.Error() == "no migration" {
			fmt.Println("  > NOTE: There is no migration.")
		} else if err != nil {
			return err, "Failed to fetch the version of the migration."
		}
		err = m.Force(int(version))
	default:
		content := "Expected the \"up\" or \"down\" or \"force\" subcommands."
		return errors.New(content), "An error occured."
	}
	if err != nil && err.Error() == "no change" {
		fmt.Println("  > NOTE: There is no change related to the operation of the migration.")
		return nil, ""
	} else if err != nil {
		return err, "Failed to establish the connection of the migration.`"
	}
	return
}