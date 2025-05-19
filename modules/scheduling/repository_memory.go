package scheduling

import (
	"booker/utils"
	"github.com/google/uuid"
)

type SchedulingMemoryRepository struct {
	data []*Scheduling
}

func NewSchedulingMemoryRepository() *SchedulingMemoryRepository {
	return &SchedulingMemoryRepository{
		data: []*Scheduling{},
	}
}

func (r *SchedulingMemoryRepository) Save(scheduling *Scheduling) error {
	r.data = append(r.data, scheduling)
	return nil
}

func (r *SchedulingMemoryRepository) List(query utils.QueryOptions) (utils.ListResponse[*Scheduling], error) {
	return utils.ListResponse[*Scheduling]{
		List:  r.data,
		Count: len(r.data),
	}, nil
}

func (r *SchedulingMemoryRepository) FindOne(id uuid.UUID) (*Scheduling, error) {
	return r.data[0], nil
}

func (r *SchedulingMemoryRepository) Update(scheduling *Scheduling) error {
	return nil
}

func (r *SchedulingMemoryRepository) Delete(id uuid.UUID) error {
	return nil
}
