package model

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/go-api-boilerplate/dto"
	"gorm.io/gorm"
)

// User is a user of the system with an account
type User struct {
	// gorm.Model with UUID
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// User fields
	Name string `gorm:"not null"`
}

// NewUserFromDTO converts a dto.User to a model.User
func NewUserFromDTO(user *dto.User) *User {
	var id uuid.UUID
	if user.ID != "" {
		id = uuid.MustParse(user.ID)
	}
	return &User{
		ID:   id,
		Name: user.Name,
	}
}

func (u *User) ToDTO() *dto.User {
	return &dto.User{
		ID:   u.ID.String(),
		Name: u.Name,
	}
}
