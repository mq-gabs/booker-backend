package scheduling

import (
	"booker/modules/base"
	"booker/utils"
	"github.com/google/uuid"
	"time"
)

type SchedulingService struct {
	repo SchedulingRepository
}

func NewSchedulingService(repo SchedulingRepository) *SchedulingService {
	return &SchedulingService{repo: repo}
}

func (s *SchedulingService) Create(dto CreateSchedulingDTO) error {
	entry := &Scheduling{
		BaseEntity: base.BaseEntity{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     dto.Name,
		Email:    dto.Email,
		Phone:    dto.Phone,
		CPF:      dto.CPF,
		Datetime: dto.Datetime,
		Status:   dto.Status,
	}
	if err := entry.Validate(); err != nil {
		return err
	}
	return s.repo.Save(entry)
}

func (s *SchedulingService) List(query utils.QueryOptions) (utils.ListResponse[*Scheduling], error) {
	return s.repo.List(query)
}

func (s *SchedulingService) FindOne(id uuid.UUID) (*Scheduling, error) {
	return s.repo.FindOne(id)
}

func (s *SchedulingService) Update(id uuid.UUID, dto UpdateSchedulingDTO) error {
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
	if dto.Phone != nil {
		existing.Phone = *dto.Phone
	}
	if dto.CPF != nil {
		existing.CPF = *dto.CPF
	}
	if dto.Datetime != nil {
		existing.Datetime = *dto.Datetime
	}
	if dto.Status != nil {
		existing.Status = *dto.Status
	}
	existing.UpdatedAt = time.Now()
	if err := existing.Validate(); err != nil {
		return err
	}
	return s.repo.Update(existing)
}

func (s *SchedulingService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
