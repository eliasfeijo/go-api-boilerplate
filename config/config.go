package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

// MigrationConfig contains the migrations configuration
type MigrationsConfig struct {
	Path string `mapstructure:"API_DATABASE_MIGRATIONS_PATH"`
}

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	Host       string           `mapstructure:"API_DATABASE_HOST"`
	Port       string           `mapstructure:"API_DATABASE_PORT"`
	User       string           `mapstructure:"API_DATABASE_USER"`
	Password   string           `mapstructure:"API_DATABASE_PASSWORD"`
	Name       string           `mapstructure:"API_DATABASE_NAME"`
	SSLMode    string           `mapstructure:"API_DATABASE_SSLMODE"`
	TimeZone   string           `mapstructure:"API_DATABASE_TIMEZONE"`
	Migrations MigrationsConfig `mapstructure:",squash"`
}

// ServerConfig contains the server configuration
type ServerConfig struct {
	Port string `mapstructure:"API_PORT"`
}

type LogConfig struct {
	Level   string `mapstructure:"API_LOG_LEVEL"`
	Verbose bool   `mapstructure:"API_LOG_VERBOSE"`
}

type AuthenticationConfig struct {
	SignatureKey string `mapstructure:"API_AUTHENTICATION_SIGNATURE_KEY"`
}

// Config is the struct that contains the API environment configuration
type Config struct {
	// Database configuration
	Database DatabaseConfig `mapstructure:",squash"`
	// Log configuration
	Log LogConfig `mapstructure:",squash"`
	// Server configuration
	Server ServerConfig `mapstructure:",squash"`
	// Authentication configuration
	Authentication AuthenticationConfig `mapstructure:",squash"`
}

var config *Config

// GetConfig returns the configuration
func GetConfig() *Config {
	if config == nil {
		config = loadConfigFromEnv(&Config{})
	}

	return config
}

// loadConfig loads the configuration from the environment variables
func loadConfigFromEnv(cfg *Config) *Config {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Database configuration
	viper.SetDefault("API_DATABASE_HOST", "")
	viper.SetDefault("API_DATABASE_PORT", "")
	viper.SetDefault("API_DATABASE_USER", "")
	viper.SetDefault("API_DATABASE_PASSWORD", "")
	viper.SetDefault("API_DATABASE_NAME", "")
	viper.SetDefault("API_DATABASE_SSLMODE", "disable")
	viper.SetDefault("API_DATABASE_TIMEZONE", "America/Sao_Paulo")
	// Migrations configuration
	viper.SetDefault("API_DATABASE_MIGRATIONS_PATH", "database/migrations")
	// Log configuration
	viper.SetDefault("API_LOG_LEVEL", "debug")
	viper.SetDefault("API_LOG_VERBOSE", "false")
	// Server configuration
	viper.SetDefault("API_PORT", "8080")
	// Authentication configuration
	viper.SetDefault("API_AUTHENTICATION_SIGNATURE_KEY", "123")

	err := viper.ReadInConfig()
	if err != nil {
		if err2, ok := err.(*os.PathError); !ok {
			err = err2
			panic(err)
		}
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}

	err = checkConfig(cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func checkConfig(cfg *Config) (err error) {
	if cfg.Database.Host == "" {
		err = errors.New("DatabaseHost is empty")
		return
	}
	if cfg.Database.Port == "" {
		err = errors.New("DatabasePort is empty")
		return
	}
	if cfg.Database.User == "" {
		err = errors.New("DatabaseUser is empty")
		return
	}
	if cfg.Database.Password == "" {
		err = errors.New("DatabasePassword is empty")
		return
	}
	if cfg.Database.Name == "" {
		err = errors.New("DatabaseName is empty")
		return
	}

	return
}
