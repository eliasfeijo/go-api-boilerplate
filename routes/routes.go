package routes

import (
	"github.com/gin-gonic/gin"
	v1 "gitlab.com/go-api-boilerplate/routes/v1"
)

func Setup(router *gin.Engine) {
	SetupHome(router.Group("/"))
	v1.SetupV1(router.Group("/v1"))
}
