package user

import (
	"booker/utils"
	"errors"

	"github.com/google/uuid"
)

type UserMemoryRepository struct {
	data []*User
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		data: []*User{},
	}
}

func (r *UserMemoryRepository) Save(user *User) error {
	r.data = append(r.data, user)
	return nil
}

func (r *UserMemoryRepository) List(query utils.QueryOptions) (utils.ListResponse[*User], error) {
	return utils.ListResponse[*User]{
		List:  r.data,
		Count: len(r.data),
	}, nil
}

func (r *UserMemoryRepository) FindOne(id uuid.UUID) (*User, error) {
	for _, u := range r.data {
		if u.ID == id {
			return u, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserMemoryRepository) Update(user *User) error {
	if _, err := r.FindOne(user.ID); err != nil {
		return err
	}

	newData := []*User{}
	for _, u := range r.data {
		if u.ID == user.ID {
			newData = append(newData, user)
		} else {
			newData = append(newData, u)
		}
	}

	r.data = newData

	return nil
}

func (r *UserMemoryRepository) Delete(id uuid.UUID) error {
	if _, err := r.FindOne(id); err != nil {
		return err
	}

	newData := []*User{}
	for _, u := range r.data {
		if u.ID != id {
			newData = append(newData, u)
		}
	}

	r.data = newData

	return nil
}
