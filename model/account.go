package model

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/go-api-boilerplate/dto"
	"gorm.io/gorm"
)

// Account contains the user's credentials
type Account struct {
	// gorm.Model with UUID
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// Account fields
	Email        string `gorm:"not null;unique"`
	PasswordHash string
	// The accounts' user
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID uuid.UUID
}

func NewAccountFromDTO(account *dto.Account) *Account {
	var id uuid.UUID
	if account.ID != "" {
		id = uuid.MustParse(account.ID)
	}
	var userID uuid.UUID
	if account.User != nil && account.User.ID != "" {
		userID = uuid.MustParse(account.User.ID)
	}
	return &Account{
		ID:           id,
		Email:        account.Email,
		PasswordHash: account.PasswordHash,
		UserID:       userID,
	}
}

func (a *Account) ToDTO() *dto.Account {
	return &dto.Account{
		ID:    a.ID.String(),
		Email: a.Email,
		User:  a.User.ToDTO(),
	}
}
