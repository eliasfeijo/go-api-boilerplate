package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/go-api-boilerplate/config"
	"gitlab.com/go-api-boilerplate/database"
)

var migrationSteps int

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long: `To run database migrations, run:

api db migrate {up|down} [flags]`,
	Run: migrate,
}

func init() {
	dbCmd.AddCommand(migrateCmd)

	cfg := config.GetConfig()

	migrationsPath := ""
	migrateCmd.PersistentFlags().StringVar(&migrationsPath, "migrations", "", "Path to migrations folder")
	migrateCmd.PersistentFlags().IntVar(&migrationSteps, "steps", 0, "Number of steps to migrate (zero means all pending migrations)")

	if migrationsPath != "" {
		cfg.Database.Migrations.Path = migrationsPath
	}
}

func migrate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalln("Please provide the direction of the migration (up, down or number of steps) eg. api db migrate up")
	}
	cfg := config.GetConfig()
	err := database.MigrateDB(args[0], cfg.Log.Verbose, cfg.Database.Migrations.Path, migrationSteps)
	if err != nil {
		log.Fatalf("Error migrating database: %s\n", err.Error())
		return
	}
}
