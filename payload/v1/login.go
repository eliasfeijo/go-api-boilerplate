package v1

import "gitlab.com/go-api-boilerplate/dto"

type LoginWithEmailAndPassword struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func LoginWithEmailAndPasswordToAccountDTO(l *LoginWithEmailAndPassword) *dto.Account {
	return &dto.Account{
		Email:    l.Email,
		Password: l.Password,
	}
}
