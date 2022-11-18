package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/go-api-boilerplate/dto"
)

// AuthorizeAccountMiddleware ...
type AuthorizeAccountMiddleware struct{}

var authorizeAccountMiddleware *AuthorizeAccountMiddleware

// AuthorizeAccount Middleware function to verify the JWT token
func AuthorizeAccount() gin.HandlerFunc {
	return authorizeAccountMiddleware.Run()
}

// NewAuthorizeAccountMiddleware ...
func NewAuthorizeAccountMiddleware() *AuthorizeAccountMiddleware {
	if authorizeAccountMiddleware == nil {
		authorizeAccountMiddleware = &AuthorizeAccountMiddleware{}
	}
	return authorizeAccountMiddleware
}

// Run ...
func (m *AuthorizeAccountMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {

		account, exists := c.Get("account")
		if !exists {
			respondWithError(c, 401, "Unauthorized")
			return
		}

		accountDTO := account.(*dto.Account)
		if c.Param("id") != accountDTO.ID {
			respondWithError(c, 403, "Forbidden")
			return
		}

		c.Next()
	}
}
