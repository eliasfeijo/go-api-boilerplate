package server

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"gitlab.com/go-api-boilerplate/config"
	_ "gitlab.com/go-api-boilerplate/docs"
	"gitlab.com/go-api-boilerplate/middleware"
	"gitlab.com/go-api-boilerplate/routes"
)

// @title          API
// @version        1.0
// @description    A basic API containing endpoints for account creation and authentication
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /v1

// @securityDefinitions.apikey Authorization Bearer Token
// @in                         header
// @name                       Authorization
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
