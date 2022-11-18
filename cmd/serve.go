package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/go-api-boilerplate/config"
	"gitlab.com/go-api-boilerplate/database"
	"gitlab.com/go-api-boilerplate/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API",
	Long: `To start the API, run:

api serve [flags]`,
	Run: serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	cfg := config.GetConfig()

	serverPort := ""

	serveCmd.PersistentFlags().StringVar(&serverPort, "port", "8080", "Port to listen on")

	if serverPort != "" {
		cfg.Server.Port = serverPort
	}

	migrationsPath := ""

	serveCmd.Flags().BoolP("migrate", "m", true, "Run migrations before starting the server")
	serveCmd.Flags().StringVar(&migrationsPath, "migrations", "", "Path to migrations folder")

	if migrationsPath != "" {
		cfg.Database.Migrations.Path = migrationsPath
	}
}

func serve(cmd *cobra.Command, args []string) {
	cfg := config.GetConfig()
	if cmd.Flag("migrate").Value.String() == "true" {
		err := database.MigrateDB("up", cfg.Log.Verbose, cfg.Database.Migrations.Path, 0)
		if err != nil {
			log.Fatalf("Error migrating database: %s\n", err.Error())
			return
		}
	}
	server.Run()
}
