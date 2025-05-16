package base

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func New() *BaseEntity {
	return &BaseEntity{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (b *BaseEntity) Validate() error {
	if b.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if b.CreatedAt.IsZero() {
		return errors.New("created_at is required")
	}
	if b.UpdatedAt.IsZero() {
		return errors.New("updated_at is required")
	}
	return nil
}
