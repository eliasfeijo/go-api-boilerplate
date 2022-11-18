package seeds

import (
	log "github.com/sirupsen/logrus"

	"gitlab.com/go-api-boilerplate/model"
	"gitlab.com/go-api-boilerplate/repository"
	"gitlab.com/go-api-boilerplate/util/crypto"
)

func (s Seed) CreateAdmin() (err error) {

	// Get the Users repository
	users := repository.GetUsers()

	count, err := users.Count(s.ctx, make(map[string]interface{}))
	if err != nil {
		return
	}

	if count > 0 {
		log.Infoln("Skipping, admin already exists")
		return
	}

	user := model.User{
		Name: "Admin",
	}

	tx := users.NewSession(s.ctx).Begin()

	// Create the user
	err = users.CreateInTransaction(s.ctx, tx, &user)
	if err != nil {
		tx.Rollback()
		return
	}

	// Get the Accounts repository
	accounts := repository.GetAccounts()

	passwordHash, err := crypto.HashPassword("admin")

	if err != nil {
		return
	}

	account := model.Account{
		Email:        "admin@admin.com",
		PasswordHash: passwordHash,
		User:         user,
	}

	tx = tx.Model(&account)

	// Create the account
	err = accounts.CreateInTransaction(s.ctx, tx, &account)

	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit().Error; err != nil {
		return
	}

	return
}
