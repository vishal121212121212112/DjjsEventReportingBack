package models

import (
	"time"
)

type GuestMaster struct {
	ID                 int       `gorm:"primaryKey;autoIncrement" json:"id"`
	GuestPrefix        string    `json:"guest_prefix"`    
	GuestFirstName     string    `json:"guest_first_name"`
	GuestMiddleName    string    `json:"guest_middle_name"`
	GuestLastName      string    `json:"guest_last_name"`
	GuestGender        string    `json:"guest_gender"`
	Designation        string    `json:"designation"`
	Organization       string    `json:"organization"`
	Email              string    `json:"email"`
	City               string    `json:"city"`
	State              string    `json:"state"`
	PersonalNumber     string    `json:"personal_number"`
	ContactNumber      string    `json:"contact_number"`
	ContactName        string    `json:"contact_name"`
	ReferenceName      string    `json:"reference_name"`
	ReferenceNumber    string    `json:"reference_number"`
	ReferenceBranchID  string    `json:"reference_branch_id"`
	CreatedOn          time.Time `json:"created_on"`
	UpdatedOn          time.Time `json:"updated_on"`
	CreatedBy          string    `json:"created_by"`
	UpdatedBy          string    `json:"updated_by"`
}

func (GuestMaster) TableName() string { return "GuestMaster" }