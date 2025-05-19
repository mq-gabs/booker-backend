package schedulingprofile

import (
	"booker/utils"
	"github.com/google/uuid"
)

type SchedulingProfileMemoryRepository struct {
	data []*SchedulingProfile
}

func NewSchedulingProfileMemoryRepository() *SchedulingProfileMemoryRepository {
	return &SchedulingProfileMemoryRepository{
		data: []*SchedulingProfile{},
	}
}

func (r *SchedulingProfileMemoryRepository) Save(scheduling *SchedulingProfile) error {
	r.data = append(r.data, scheduling)
	return nil
}

func (r *SchedulingProfileMemoryRepository) List(query utils.QueryOptions) (utils.ListResponse[*SchedulingProfile], error) {
	return utils.ListResponse[*SchedulingProfile]{
		List:  r.data,
		Count: len(r.data),
	}, nil
}

func (r *SchedulingProfileMemoryRepository) FindOne(id uuid.UUID) (*SchedulingProfile, error) {
	return r.data[0], nil
}

func (r *SchedulingProfileMemoryRepository) Update(scheduling *SchedulingProfile) error {
	return nil
}

func (r *SchedulingProfileMemoryRepository) Delete(id uuid.UUID) error {
	return nil
}
