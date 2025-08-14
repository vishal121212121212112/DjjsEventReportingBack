package models

import (
	"time"
)

type EventHistory struct {
	ID                       int       `json:"id"`
	FKProgramMasterID        int       `json:"fk_program_master_id"`
	FKBranchID               int       `json:"fk_branch_id"`
	ProgramOrganisedBy       string    `json:"program_organised_by"`
	Scale                    string    `json:"scale"`
	StartDate                time.Time `json:"start_date"`
	EndDate                  time.Time `json:"end_date"`
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
	ThemePurposeMessageGiven string    `json:"theme_purpose_message_given"`
	IsComplete               bool      `json:"is_complete"` 
	Steps                    int       `json:"steps"` 
	SpiritualOrator          string    `json:"spiritual_orator"`
	Language                 string    `json:"language"`
	NoOfBeneficiaries        int       `json:"no_of_beneficiaries"`
	NoOfInitiationMale       int       `json:"no_of_initiation_male"`
	NoOfInitiationFemale     int       `json:"no_of_initiation_female"`
	NoOfInitiationKids       int       `json:"no_of_initiation_kids"`
	TimeDuration             string    `json:"time_duration"`
	ProgramVenue             string    `json:"program_venue"`
	City                     string    `json:"city"`
	State                    string    `json:"state"`
	Country                  string    `json:"country"`
	PINCode                  string    `json:"pin_code"`
	Frequency                string    `json:"frequency"`
	ProcRouteStart           string    `json:"procession_route_start"`
	ProcRouteEnd             string    `json:"procession_route_end"`
	NoOfParticipants         int       `json:"no_of_participants"`
	EstimatedAreaCoveredKm2  float64   `json:"estimated_area_covered_km2"`
	CreatedOn                time.Time `json:"created_on"`
	UpdatedOn                time.Time `json:"updated_on"`
	CreatedBy                string    `json:"created_by"`
	UpdatedBy                string    `json:"updated_by"`
}

func (EventHistory) TableName() string { return "EventHistory" }