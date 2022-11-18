package v1

import (
	"github.com/gin-gonic/gin"
	controller "gitlab.com/go-api-boilerplate/controller/v1"
	"gitlab.com/go-api-boilerplate/middleware"
)

func SetupAccounts(router *gin.RouterGroup) {
	accounts := controller.NewAccounts()
	router.POST("/login", accounts.Login())
	router.POST("/", accounts.CreateAccount())
	router.PUT("/:id",
		middleware.AuthenticateJwt(),
		middleware.AuthorizeAccount(),
		accounts.UpdateAccount())
	router.DELETE("/:id",
		middleware.AuthenticateJwt(),
		middleware.AuthorizeAccount(),
		accounts.DeleteAccount())
}
