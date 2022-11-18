package v1

import "gitlab.com/go-api-boilerplate/dto"

type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func UserToDTO(user *User) *dto.User {
	return &dto.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
