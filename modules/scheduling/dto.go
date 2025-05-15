package scheduling

import (
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
