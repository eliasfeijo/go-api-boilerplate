package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"gitlab.com/go-api-boilerplate/config"
	"gitlab.com/go-api-boilerplate/middleware"
	"gitlab.com/go-api-boilerplate/routes"
)

func Run() (err error) {
	cfg := config.GetConfig()

	gin.DefaultWriter = log.StandardLogger().WriterLevel(log.DebugLevel)
	gin.DefaultErrorWriter = log.StandardLogger().WriterLevel(log.ErrorLevel)

	router := gin.New()
	router.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery())

	middleware.InitializeMiddlewares()

	// Setup routes
	routes.Setup(router)

	router.Run(":" + cfg.Server.Port)

	return
}
