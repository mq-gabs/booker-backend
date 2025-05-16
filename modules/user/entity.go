package user

import (
	"booker/modules/base"
	"errors"
	"net/mail"
)

type User struct {
	base.BaseEntity
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New() *User {
	b := base.New()
	return &User{
		BaseEntity: *b,
	}
}

func (u *User) Validate() error {
	if err := u.BaseEntity.Validate(); err != nil {
		return err
	}
	if u.Name == "" {
		return errors.New("name is required")
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email")
	}
	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}
