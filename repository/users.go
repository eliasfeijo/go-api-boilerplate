package repository

import (
	"gitlab.com/go-api-boilerplate/database"
	"gitlab.com/go-api-boilerplate/model"
)

type Users interface {
	Repository[model.User]
}

var usersInstance Users

func GetUsers() Users {
	if usersInstance == nil {
		db := database.GetConn()
		usersInstance = &repository[model.User]{
			db: db,
		}
	}

	return usersInstance
}
