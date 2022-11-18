package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	v1 "gitlab.com/go-api-boilerplate/routes/v1"
)

func Setup(router *gin.Engine) {
	SetupHome(router.Group("/"))
	v1.SetupV1(router.Group("/v1"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
