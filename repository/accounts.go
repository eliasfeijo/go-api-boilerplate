package repository

import (
	"gitlab.com/go-api-boilerplate/database"
	"gitlab.com/go-api-boilerplate/model"
)

type Accounts interface {
	Repository[model.Account]
}

var accountsInstance Accounts

func GetAccounts() Accounts {
	if accountsInstance == nil {
		db := database.GetConn()
		accountsInstance = &repository[model.Account]{
			db: db,
		}
	}

	return accountsInstance
}
