package v1

import "gitlab.com/go-api-boilerplate/dto"

type Account struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	User  *User  `json:"user,omitempty"`
}

func NewAccountFromDTO(account *dto.Account) *Account {
	var user *User
	if account.User != nil {
		user = NewUserFromDTO(account.User)
	}
	return &Account{
		ID:    account.ID,
		Email: account.Email,
		User:  user,
	}
}
