package v1

import (
	"github.com/gin-gonic/gin"
	payload "gitlab.com/go-api-boilerplate/payload/v1"
	response "gitlab.com/go-api-boilerplate/response/v1"
	"gitlab.com/go-api-boilerplate/service"
)

type Accounts interface {
	Login() gin.HandlerFunc
	CreateAccount() gin.HandlerFunc
	UpdateAccount() gin.HandlerFunc
	DeleteAccount() gin.HandlerFunc
}

type accounts struct {
	as service.Accounts
}

func NewAccounts() Accounts {
	return &accounts{
		as: service.NewAccounts(),
	}
}

func (a accounts) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		var p payload.Account
		err := c.BindJSON(&p)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		account := payload.AccountToDTO(&p)

		jwt, err := a.as.Login(c, account)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		if account == nil {
			c.JSON(401, gin.H{"error": "invalid user"})
			return
		}

		r := gin.H{
			"jwt":     jwt,
			"account": response.NewAccountFromDTO(account),
		}
		c.JSON(200, r)
	}
}

func (a accounts) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {

		var p payload.Account
		err := c.BindJSON(&p)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		account := payload.AccountToDTO(&p)

		jwt, err := a.as.CreateAccount(c, account)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		r := gin.H{
			"jwt":     jwt,
			"account": response.NewAccountFromDTO(account),
		}
		c.JSON(200, r)
	}
}

func (a accounts) UpdateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {

		var p payload.Account
		err := c.BindJSON(&p)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		account := payload.AccountToDTO(&p)
		account.ID = c.Param("id")

		err = a.as.UpdateAccount(c, account)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, response.NewAccountFromDTO(account))
	}
}

func (a accounts) DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		deleted, err := a.as.DeleteAccount(c, id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"deleted": deleted})
	}
}
