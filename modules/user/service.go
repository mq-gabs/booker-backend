package user

import (
	"booker/modules/base"
	"booker/utils"
	"github.com/google/uuid"
	"time"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(dto CreateUserDTO) error {
	user := &User{
		BaseEntity: base.BaseEntity{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
	if err := user.Validate(); err != nil {
		return err
	}
	return s.repo.Save(user)
}

func (s *UserService) List(query utils.QueryOptions) (utils.ListResponse[*User], error) {
	return s.repo.List(query)
}

func (s *UserService) FindOne(id uuid.UUID) (*User, error) {
	return s.repo.FindOne(id)
}

func (s *UserService) Update(id uuid.UUID, dto UpdateUserDTO) error {
	existing, err := s.repo.FindOne(id)
	if err != nil {
		return err
	}
	if dto.Name != nil {
		existing.Name = *dto.Name
	}
	if dto.Email != nil {
		existing.Email = *dto.Email
	}
	if dto.Password != nil {
		existing.Password = *dto.Password
	}
	existing.UpdatedAt = time.Now()
	if err := existing.Validate(); err != nil {
		return err
	}
	return s.repo.Update(existing)
}

func (s *UserService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
