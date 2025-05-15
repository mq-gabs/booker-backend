package schedulingprofile

import (
	"booker/utils"
	"github.com/google/uuid"
)

type SchedulingProfileRepository interface {
	Save(profile *SchedulingProfile) error
	List(query utils.QueryOptions) (utils.ListResponse[*SchedulingProfile], error)
	FindOne(id uuid.UUID) (*SchedulingProfile, error)
	Update(profile *SchedulingProfile) error
	Delete(id uuid.UUID) error
}
