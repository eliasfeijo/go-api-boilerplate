package middleware

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/go-api-boilerplate/config"
	payload "gitlab.com/go-api-boilerplate/payload/v1"
)

// AuthenticateJwtMiddleware ...
type AuthenticateJwtMiddleware struct {
	signatureKey jwk.Key
}

var authenticateJwtMiddleware *AuthenticateJwtMiddleware

// AuthenticateJwt Middleware function to verify the JWT token
func AuthenticateJwt() gin.HandlerFunc {
	return authenticateJwtMiddleware.Run()
}

// NewAuthenticateJwtMiddleware ...
func NewAuthenticateJwtMiddleware() *AuthenticateJwtMiddleware {
	if authenticateJwtMiddleware == nil {
		cfg := config.GetConfig()

		key, err := jwk.FromRaw([]byte(cfg.Authentication.SignatureKey))

		if err != nil {
			log.Fatalf("Error creating signature key: %s\n", err.Error())
			return nil
		}

		authenticateJwtMiddleware = &AuthenticateJwtMiddleware{
			signatureKey: key,
		}
	}

	return authenticateJwtMiddleware
}

// Run ...
func (m *AuthenticateJwtMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, statusCode, errMsg := m.extractTokenFromHeader(c)

		if errMsg != "" {
			respondWithError(c, statusCode, errMsg)
			return
		}

		invalidJWTError := "Invalid JWT token"
		defaultStatusCode := 401

		parsed, internalErr := jwt.Parse([]byte(token), jwt.WithKey(jwa.HS256, m.signatureKey))
		if internalErr != nil {
			respondWithError(c, defaultStatusCode, invalidJWTError)
			return
		}

		a, exists := parsed.Get("account")
		if !exists {
			respondWithError(c, defaultStatusCode, invalidJWTError)
			return
		}

		p := payload.Account{}
		err := json.Unmarshal([]byte(a.(string)), &p)
		if err != nil {
			respondWithError(c, defaultStatusCode, invalidJWTError)
			return
		}

		log.Printf("p: %v", p)
		account := payload.AccountToDTO(&p)
		log.Printf("account: %v", account)

		c.Set("account", account)
	}
}

func (m *AuthenticateJwtMiddleware) extractTokenFromHeader(c *gin.Context) (string, int, string) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", 401, "Missing Authorization Header"
	}
	splitHeader := strings.Split(authorizationHeader, "Bearer")
	if len(splitHeader) != 2 {
		return "", 401, "Invalid Authorization Header"
	}
	return strings.TrimSpace(splitHeader[1]), 0, ""
}
