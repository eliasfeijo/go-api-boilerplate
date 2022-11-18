package v1

import "gitlab.com/go-api-boilerplate/dto"

type User struct {
	Name string `json:"name,omitempty"`
}

func UserToDTO(user *User) *dto.User {
	return &dto.User{
		Name: user.Name,
	}
}
