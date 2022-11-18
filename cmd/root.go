package cmd

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/go-api-boilerplate/config"
	"gitlab.com/go-api-boilerplate/database"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "API root command",
	Long: `To start the API, run:

api serve`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	cfg := config.GetConfig()

	verbose := false
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose logging")
	if verbose {
		cfg.Log.Verbose = verbose
	}

	level := ""
	rootCmd.PersistentFlags().StringVarP(&level, "log", "", "debug", "Log level")
	if level != "" {
		cfg.Log.Level = level
	}
}

func initConfig() {
	cfg := config.GetConfig()
	setupLog(cfg)
	database.GetConn()
}

func setupLog(cfg *config.Config) {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	if cfg.Log.Verbose {
		log.SetLevel(log.TraceLevel)
		log.Infoln("Verbose logging enabled")
		return
	}
	switch strings.ToLower(cfg.Log.Level) {
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	}
}
