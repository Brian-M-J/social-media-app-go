package users

import (
	"context"

	"github.com/Brian-M-J/social-media-app-go/internals/dto"
	"github.com/Brian-M-J/social-media-app-go/models/users"
)

type User struct {
	User *dto.User
}

func New() *User {
	return &User{}
}

func (u *User) Create(ctx context.Context) {
	m := users.New()
	m.Name = u.User.Name
	m.Email = u.User.Email
	m.Password = u.User.Password
	m.Create(ctx)
	u.User.CreatedAt = &m.CreatedAt
	u.User.UpdatedAt = nil
}

func (u *User) Get(ctx context.Context) error {
	m := users.New()
	m.ID = u.User.ID
	m.User = u.User
	if err := m.Get(ctx); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(ctx context.Context) error {
	m := users.New()
	m.ID = u.User.ID
	m.User = u.User
	if err := m.Delete(ctx); err != nil {
		return err
	}
	return nil
}
