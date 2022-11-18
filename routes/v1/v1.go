package v1

import (
	"github.com/gin-gonic/gin"
)

func SetupV1(router *gin.RouterGroup) {
	SetupAccounts(router.Group("/accounts"))
}
