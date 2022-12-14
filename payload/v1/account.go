package v1

import "gitlab.com/go-api-boilerplate/dto"

type Account struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	User     *User  `json:"user,omitempty"`
}

type AccountWithID struct {
	Account
	ID string `json:"id,omitempty"`
}

func AccountToDTO(account *Account) *dto.Account {
	var user *dto.User
	if account.User != nil {
		user = UserToDTO(account.User)
	}
	return &dto.Account{
		Email:    account.Email,
		Password: account.Password,
		User:     user,
	}
}

func AccountWithIDToDTO(account *AccountWithID) *dto.Account {
	var user *dto.User
	if account.User != nil {
		user = UserToDTO(account.User)
	}
	return &dto.Account{
		ID:       account.ID,
		Email:    account.Email,
		Password: account.Password,
		User:     user,
	}
}
