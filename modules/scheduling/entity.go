package scheduling

import (
	"booker/modules/base"
	"errors"
	"net/mail"
	"strings"
	"time"
)

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
