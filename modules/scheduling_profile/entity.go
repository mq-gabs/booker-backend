package schedulingprofile

import (
	"booker/modules/base"
	"errors"
	"strings"
	"time"
)

type SchedulingHour struct {
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

type SchedulingProfile struct {
	base.BaseEntity
	Name      string         `json:"name"`
	Monday    SchedulingHour `json:"monday"`
	Tuesday   SchedulingHour `json:"tuesday"`
	Wednesday SchedulingHour `json:"wednesday"`
	Thursday  SchedulingHour `json:"thursday"`
	Friday    SchedulingHour `json:"friday"`
	Saturday  SchedulingHour `json:"saturday"`
	Sunday    SchedulingHour `json:"sunday"`
}

func (sh *SchedulingHour) Validate() error {
	if sh.EndAt.Before(sh.StartAt) {
		return errors.New("end time must be after start time")
	}
	return nil
}

func (sp *SchedulingProfile) Validate() error {
	if err := sp.BaseEntity.Validate(); err != nil {
		return err
	}
	if strings.TrimSpace(sp.Name) == "" {
		return errors.New("name is required")
	}
	days := []SchedulingHour{
		sp.Monday, sp.Tuesday, sp.Wednesday, sp.Thursday, sp.Friday, sp.Saturday, sp.Sunday,
	}
	for _, day := range days {
		if err := day.Validate(); err != nil {
			return err
		}
	}
	return nil
}
