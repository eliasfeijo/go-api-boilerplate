package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/go-api-boilerplate/config"
	"gitlab.com/go-api-boilerplate/database"
	"gitlab.com/go-api-boilerplate/database/seeds"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Runs database migrations and seeds",
	Long: `To run database migrations and seeds, run:

api db setup [flags]`,
	Run: setup,
}

func init() {
	dbCmd.AddCommand(setupCmd)

	cfg := config.GetConfig()

	migrationsPath := ""
	setupCmd.PersistentFlags().StringVar(&migrationsPath, "migrations", "", "Path to migrations folder")

	if migrationsPath != "" {
		cfg.Database.Migrations.Path = migrationsPath
	}
}

func setup(cmd *cobra.Command, args []string) {
	cfg := config.GetConfig()
	err := database.MigrateDB("up", cfg.Log.Verbose, cfg.Database.Migrations.Path, 0)
	if err != nil {
		log.Fatalf("Error migrating database: %s\n", err.Error())
		return
	}
	err = seeds.SeedDB()
	if err != nil {
		log.Fatalf("Error seeding database: %s\n", err.Error())
		return
	}
}
