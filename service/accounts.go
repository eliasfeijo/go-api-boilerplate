package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/go-api-boilerplate/config"
	"gitlab.com/go-api-boilerplate/dto"
	"gitlab.com/go-api-boilerplate/model"
	"gitlab.com/go-api-boilerplate/repository"
	response "gitlab.com/go-api-boilerplate/response/v1"
	"gitlab.com/go-api-boilerplate/util/crypto"
)

type Accounts interface {
	Login(ctx context.Context, account *dto.Account) (string, error)
	GetAccount(ctx context.Context, id string) (*dto.Account, error)
	CreateAccount(ctx context.Context, account *dto.Account) (string, error)
	UpdateAccount(ctx context.Context, account *dto.Account) error
	DeleteAccount(ctx context.Context, id string) (bool, error)
}

type accounts struct {
	ra              repository.Accounts
	ru              repository.Users
	jwtSignatureKey jwk.Key
}

func NewAccounts() Accounts {
	cfg := config.GetConfig()

	key, err := jwk.FromRaw([]byte(cfg.Authentication.SignatureKey))
	if err != nil {
		log.Fatalf("Error creating signature key: %s\n", err.Error())
		return nil
	}

	return &accounts{
		ra:              repository.GetAccounts(),
		ru:              repository.GetUsers(),
		jwtSignatureKey: key,
	}
}

func (a *accounts) Login(ctx context.Context, account *dto.Account) (token string, err error) {

	tx := a.ra.NewSession(ctx).Begin().Preload("User")

	user, err := a.ra.FindInTransaction(ctx, tx, map[string]interface{}{"email": account.Email}, nil)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	errUnauthorized := fmt.Errorf("unauthorized")

	if len(user) == 0 {
		err = errUnauthorized
		return
	}

	passwordMatch := crypto.CheckPasswordHash(account.Password, user[0].PasswordHash)
	if !passwordMatch {
		err = errUnauthorized
		return
	}

	account.ID = user[0].ID.String()
	account.Email = user[0].Email
	account.Password = ""
	account.PasswordHash = ""
	account.User = user[0].User.ToDTO()

	r := response.NewAccountFromDTO(account)

	j, err := json.Marshal(r)
	if err != nil {
		return
	}

	t, err := jwt.NewBuilder().
		Claim("account", string(j)).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(time.Hour * 1)).
		Build()
	if err != nil {
		return
	}

	signed, err := jwt.Sign(t, jwt.WithKey(jwa.HS256, a.jwtSignatureKey))
	if err != nil {
		return
	}

	token = string(signed)

	return
}

func (a *accounts) GetAccount(ctx context.Context, id string) (*dto.Account, error) {
	entity, err := a.ra.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return entity.ToDTO(), nil
}

func (a *accounts) CreateAccount(ctx context.Context, account *dto.Account) (token string, err error) {

	tx := a.ru.NewSession(ctx).Begin()

	user := model.NewUserFromDTO(account.User)
	err = a.ru.CreateInTransaction(ctx, tx, user)
	if err != nil {
		tx.Rollback()
		return
	}

	account.ID = uuid.New().String()
	account.PasswordHash, _ = crypto.HashPassword(account.Password)
	entity := model.NewAccountFromDTO(account)
	entity.UserID = user.ID

	tx = tx.Model(&entity).Preload("User")
	err = a.ra.CreateInTransaction(ctx, tx, entity)
	if err != nil {
		tx.Rollback()
		return
	}
	account.ID = entity.ID.String()
	account.Password = ""
	account.PasswordHash = ""

	if err = tx.Commit().Error; err != nil {
		return
	}

	r := response.NewAccountFromDTO(account)

	j, err := json.Marshal(r)
	if err != nil {
		return
	}

	t, err := jwt.NewBuilder().
		Claim("account", string(j)).
		IssuedAt(time.Now()).
		Expiration(time.Now().Add(time.Hour * 1)).
		Build()
	if err != nil {
		return
	}

	signed, err := jwt.Sign(t, jwt.WithKey(jwa.HS256, a.jwtSignatureKey))
	if err != nil {
		return
	}

	token = string(signed)

	return
}

func (a *accounts) UpdateAccount(ctx context.Context, account *dto.Account) (err error) {
	entity := model.NewAccountFromDTO(account)
	tx := a.ra.NewSession(ctx).Model(&entity).Begin()

	if account.User != nil {
		user := model.NewUserFromDTO(account.User)
		err = a.ru.UpdateInTransaction(ctx, tx, user)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	if account.Password != "" {
		account.PasswordHash, _ = crypto.HashPassword(account.Password)
	}

	err = a.ra.UpdateInTransaction(ctx, tx, entity)
	if err != nil {
		tx.Rollback()
		return
	}

	account.Email = entity.Email
	account.Password = ""
	account.PasswordHash = ""

	err = tx.Commit().Error
	return
}

func (a *accounts) DeleteAccount(ctx context.Context, id string) (bool, error) {
	return a.ra.Delete(ctx, id)
}
