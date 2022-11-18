package v1

import "gitlab.com/go-api-boilerplate/dto"

type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewUserFromDTO(user *dto.User) *User {
	return &User{
		ID:   user.ID,
		Name: user.Name,
	}
}
