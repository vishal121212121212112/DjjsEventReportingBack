package models

import "time"

type ProgramVolunteer struct {
	ID                     int       `json:"id"`                     
	FKProgramID            int       `json:"fk_program_id"`               
	FKBranchID             int       `json:"fk_branch_id"`              
	MemberID               int       `json:"member_id"`                   
	VolunteerName          string    `json:"volunteer_name"`             
	Gender                 string    `json:"gender"`                      
	VolunteerBranchName    string    `json:"volunteer_branch_name"`      
	ContactNumber          string    `json:"contact_number"`             
	NoOfDaysInvolvedInSewa int       `json:"no_of_days_involved_in_sewa"` 
	PermanentSewa          string    `json:"permanent_sewa"`              
	SewaDepartment         string    `json:"sewa_department"`           
	CreatedOn              time.Time `json:"created_on"`             
	UpdatedOn              time.Time `json:"updated_on"`
	CreatedBy              string    `json:"created_by"`
	UpdatedBy              string    `json:"updated_by"`
}

func (ProgramVolunteer) TableName() string { return "ProgramVolunteer" }