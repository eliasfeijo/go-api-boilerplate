package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/go-api-boilerplate/database/seeds"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run database seeds",
	Long: `To run database seeds, run:

api db seed	[names...]`,
	Run: seed,
}

func init() {
	dbCmd.AddCommand(seedCmd)
}

func seed(cmd *cobra.Command, args []string) {
	err := seeds.SeedDB(args...)
	if err != nil {
		log.Fatalf("Error seeding database: %s\n", err.Error())
		return
	}
}
