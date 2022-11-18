package cmd

import (
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Run database migrations and seeds",
	Long: `To run database migrations and seeds, run:

api db {setup|migrate|seed} [flags]`,
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
