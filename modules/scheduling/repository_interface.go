package scheduling

import (
	"booker/utils"
	"github.com/google/uuid"
)

type SchedulingRepository interface {
	Save(scheduling *Scheduling) error
	List(query utils.QueryOptions) (utils.ListResponse[*Scheduling], error)
	FindOne(id uuid.UUID) (*Scheduling, error)
	Update(scheduling *Scheduling) error
	Delete(id uuid.UUID) error
}
