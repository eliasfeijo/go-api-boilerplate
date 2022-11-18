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

// Login godoc
// @Summary     Authenticates an account
// @Description returns a JWT and account information
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Param       Account body     payload.LoginWithEmailAndPassword true "Login body"
// @Success     200     {object} response.AccountWithJWT
// @Failure     400     {object} response.Error
// @Failure     404     {object} response.Error
// @Failure     500     {object} response.Error
// @Router      /accounts/login [post]
func (a accounts) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		var p payload.LoginWithEmailAndPassword
		err := c.BindJSON(&p)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		account := payload.LoginWithEmailAndPasswordToAccountDTO(&p)

		jwt, err := a.as.Login(c, account)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		if account == nil {
			c.JSON(401, gin.H{"error": "invalid user"})
			return
		}

		c.JSON(200, response.AccountWithJWT{
			JWT:     jwt,
			Account: response.NewAccountFromDTO(account),
		})
	}
}

// Create godoc
// @Summary     Creates an account
// @Description returns a JWT and account information
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Param       Account body     payload.Account true "Account body"
// @Success     200     {object} response.AccountWithJWT
// @Failure     400     {object} response.Error
// @Failure     404     {object} response.Error
// @Failure     500     {object} response.Error
// @Router      /accounts [post]
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

		c.JSON(200, response.AccountWithJWT{
			JWT:     jwt,
			Account: response.NewAccountFromDTO(account),
		})
	}
}

// Update godoc
// @Summary     Updates an account and user
// @Description returns the updated account information
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Param       id      path     string          true "Account ID"
// @Param       Account body     payload.Account true "Account body"
// @Success     200     {object} response.Account
// @Failure     400     {object} response.Error
// @Failure     404     {object} response.Error
// @Failure     500     {object} response.Error
// @Router      /accounts/{id} [put]
// @Security    Authorization Bearer Token
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

// Delete godoc
// @Summary     Deletes an account and user
// @Description returns 204 on success
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Account ID"
// @Success     204 {object} nil
// @Failure     404 {object} response.Error
// @Failure     500 {object} response.Error
// @Router      /accounts/{id} [delete]
// @Security    Authorization Bearer Token
func (a accounts) DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		deleted, err := a.as.DeleteAccount(c, id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, gin.H{"deleted": deleted})
	}
}
