package database

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/go-api-boilerplate/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetConn returns the database connection
func GetConn() *gorm.DB {
	if db == nil {
		db = connect()
	}
	return db
}

// connect connects to the database
func connect() *gorm.DB {

	cfg := config.GetConfig()

	dsn := "host=" + cfg.Database.Host + " user=" + cfg.Database.User + " password=" + cfg.Database.Password + " dbname=" + cfg.Database.Name + " port=" + cfg.Database.Port + " sslmode=" + cfg.Database.SSLMode + " TimeZone=" + cfg.Database.TimeZone

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = dbConn.ConnPool.QueryContext(ctx, "SELECT 1")
	if err != nil {
		log.Fatal(err)
	}

	log.Debugln("Connected to database")

	return dbConn
}
