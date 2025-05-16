package schedulingprofile

type CreateSchedulingProfileDTO struct {
	Name      string          `json:"name"`
	Monday    *SchedulingHour `json:"monday"`
	Tuesday   *SchedulingHour `json:"tuesday"`
	Wednesday *SchedulingHour `json:"wednesday"`
	Thursday  *SchedulingHour `json:"thursday"`
	Friday    *SchedulingHour `json:"friday"`
	Saturday  *SchedulingHour `json:"saturday"`
	Sunday    *SchedulingHour `json:"sunday"`
}

type UpdateSchedulingProfileDTO struct {
	Name      *string         `json:"name,omitempty"`
	Monday    *SchedulingHour `json:"monday,omitempty"`
	Tuesday   *SchedulingHour `json:"tuesday,omitempty"`
	Wednesday *SchedulingHour `json:"wednesday,omitempty"`
	Thursday  *SchedulingHour `json:"thursday,omitempty"`
	Friday    *SchedulingHour `json:"friday,omitempty"`
	Saturday  *SchedulingHour `json:"saturday,omitempty"`
	Sunday    *SchedulingHour `json:"sunday,omitempty"`
}
