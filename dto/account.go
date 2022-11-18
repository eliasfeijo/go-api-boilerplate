package dto

type Account struct {
	ID           string
	Email        string
	Password     string
	PasswordHash string
	User         *User
}
