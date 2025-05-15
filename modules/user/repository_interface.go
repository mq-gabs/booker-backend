package user

import (
	"booker/utils"
	"github.com/google/uuid"
)

type UserRepository interface {
	Save(user *User) error
	List(query utils.QueryOptions) (utils.ListResponse[*User], error)
	FindOne(id uuid.UUID) (*User, error)
	Update(user *User) error
	Delete(id uuid.UUID) error
}
