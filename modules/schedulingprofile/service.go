package schedulingprofile

import (
	"booker/modules/base"
	"booker/utils"
	"github.com/google/uuid"
	"time"
)

type SchedulingProfileService struct {
	repo SchedulingProfileRepository
}

func NewSchedulingProfileService(repo SchedulingProfileRepository) *SchedulingProfileService {
	return &SchedulingProfileService{repo: repo}
}

func (s *SchedulingProfileService) Create(dto CreateSchedulingProfileDTO) error {
	profile := &SchedulingProfile{
		BaseEntity: base.BaseEntity{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:      dto.Name,
		Monday:    dto.Monday,
		Tuesday:   dto.Tuesday,
		Wednesday: dto.Wednesday,
		Thursday:  dto.Thursday,
		Friday:    dto.Friday,
		Saturday:  dto.Saturday,
		Sunday:    dto.Sunday,
	}
	if err := profile.Validate(); err != nil {
		return err
	}
	return s.repo.Save(profile)
}

func (s *SchedulingProfileService) List(query utils.QueryOptions) (utils.ListResponse[*SchedulingProfile], error) {
	return s.repo.List(query)
}

func (s *SchedulingProfileService) FindOne(id uuid.UUID) (*SchedulingProfile, error) {
	return s.repo.FindOne(id)
}

func (s *SchedulingProfileService) Update(id uuid.UUID, dto UpdateSchedulingProfileDTO) error {
	existing, err := s.repo.FindOne(id)
	if err != nil {
		return err
	}
	if dto.Name != nil {
		existing.Name = *dto.Name
	}
	if dto.Monday != nil {
		existing.Monday = dto.Monday
	}
	if dto.Tuesday != nil {
		existing.Tuesday = dto.Tuesday
	}
	if dto.Wednesday != nil {
		existing.Wednesday = dto.Wednesday
	}
	if dto.Thursday != nil {
		existing.Thursday = dto.Thursday
	}
	if dto.Friday != nil {
		existing.Friday = dto.Friday
	}
	if dto.Saturday != nil {
		existing.Saturday = dto.Saturday
	}
	if dto.Sunday != nil {
		existing.Sunday = dto.Sunday
	}
	existing.UpdatedAt = time.Now()
	if err := existing.Validate(); err != nil {
		return err
	}
	return s.repo.Update(existing)
}

func (s *SchedulingProfileService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
