package service

import (
	"context"

	"gitlab.com/go-api-boilerplate/dto"
	"gitlab.com/go-api-boilerplate/model"
	"gitlab.com/go-api-boilerplate/repository"
)

type Users interface {
	GetUsers(ctx context.Context, filter interface{}, options *FindOptions) ([]*dto.User, error)
	GetUser(ctx context.Context, id string) (*dto.User, error)
	CreateUser(ctx context.Context, user *dto.User) error
	UpdateUser(ctx context.Context, user *dto.User) error
	DeleteUser(ctx context.Context, id string) (bool, error)
}

type users struct {
	repository repository.Users
}

func NewUsers() Users {
	return &users{
		repository: repository.GetUsers(),
	}
}

func (u *users) GetUsers(ctx context.Context, filter interface{}, options *FindOptions) ([]*dto.User, error) {
	opts := &repository.FindOptions{
		Limit: options.Limit,
		Page:  options.Page,
	}
	entities, err := u.repository.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	users := make([]*dto.User, len(entities))
	for i, entity := range entities {
		users[i] = entity.ToDTO()
	}

	return users, nil
}

func (u *users) GetUser(ctx context.Context, id string) (*dto.User, error) {
	entity, err := u.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return entity.ToDTO(), nil
}

func (u *users) CreateUser(ctx context.Context, user *dto.User) error {
	entity := model.NewUserFromDTO(user)
	err := u.repository.Create(ctx, entity)
	if err != nil {
		return err
	}
	user.ID = entity.ID.String()
	return nil
}

func (u *users) UpdateUser(ctx context.Context, user *dto.User) error {
	entity := model.NewUserFromDTO(user)
	err := u.repository.Update(ctx, entity)
	user.Name = entity.Name
	return err
}

func (u *users) DeleteUser(ctx context.Context, id string) (bool, error) {
	return u.repository.Delete(ctx, id)
}
