package scheduling

import (
	"booker/modules/base"
	"errors"
	"net/mail"
	"strings"
	"time"
)

type CreateSchedulingDTO struct {
	Name     string           `json:"name"`
	Email    string           `json:"email"`
	Phone    string           `json:"phone"`
	CPF      string           `json:"cpf"`
	Datetime time.Time        `json:"datetime"`
	Status   SchedulingStatus `json:"status"`
}

type UpdateSchedulingDTO struct {
	Name     *string           `json:"name,omitempty"`
	Email    *string           `json:"email,omitempty"`
	Phone    *string           `json:"phone,omitempty"`
	CPF      *string           `json:"cpf,omitempty"`
	Datetime *time.Time        `json:"datetime,omitempty"`
	Status   *SchedulingStatus `json:"status,omitempty"`
}

type SchedulingStatus string

const (
	Scheduled SchedulingStatus = "SCHEDULED"
	Cancelled SchedulingStatus = "CANCELLED"
	Done      SchedulingStatus = "DONE"
)

type Scheduling struct {
	base.BaseEntity
	Name     string           `json:"name"`
	Email    string           `json:"email"`
	Phone    string           `json:"phone"`
	CPF      string           `json:"cpf"`
	Datetime time.Time        `json:"datetime"`
	Status   SchedulingStatus `json:"status"`
}

func (s *Scheduling) Validate() error {
	if err := s.BaseEntity.Validate(); err != nil {
		return err
	}
	if strings.TrimSpace(s.Name) == "" {
		return errors.New("name is required")
	}
	if _, err := mail.ParseAddress(s.Email); err != nil {
		return errors.New("invalid email")
	}
	if s.Phone == "" {
		return errors.New("phone is required")
	}
	if s.CPF == "" {
		return errors.New("CPF is required")
	}
	if s.Status != Scheduled && s.Status != Cancelled && s.Status != Done {
		return errors.New("invalid scheduling status")
	}
	return nil
}
