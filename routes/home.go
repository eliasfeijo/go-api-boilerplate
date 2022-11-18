package routes

import "github.com/gin-gonic/gin"

func SetupHome(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
}
