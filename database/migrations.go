package database

import (
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

// MigrationLogger is a custom logger for the migrations
type MigrationLogger struct {
	verbose bool
}

// Printf implementation of the migrate.Logger interface
func (l *MigrationLogger) Printf(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Verbose implementation of the migrate.Logger interface
func (l *MigrationLogger) Verbose() bool {
	return l.verbose
}

var logger MigrationLogger

// MigrateDB runs the database migrations
// cmd can be "up", "down" or an integer with the steps to run
func MigrateDB(cmd string, verbose bool, path string, steps int) (err error) {
	log.Infoln("Running migrations")

	database, err := db.DB()
	if err != nil {
		return
	}
	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		"postgres", driver)
	if err != nil {
		return
	}

	logger = MigrationLogger{verbose: verbose}
	m.Log = &logger

	switch strings.ToUpper(cmd) {
	case "UP":
		if steps == 0 {
			err = m.Up()
		} else {
			err = m.Steps(steps)
		}
	case "DOWN":
		if steps == 0 {
			err = m.Down()
		} else {
			err = m.Steps(-steps)
		}
	}
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Infoln("No change")
			return nil
		}
		return
	}
	log.Infoln("Migrations ran successfully")
	return nil
}
