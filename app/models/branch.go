package models

import (
	"github.com/google/uuid"
)

type Branch struct {
	ID               uuid.UUID `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	BranchName       string    `json:"branch_name" gorm:"column:branch_name"`
	ParentBranch     string    `json:"parent_branch" gorm:"column:parent_branch"`
	CoordinatorName  string    `json:"coordinator_name" gorm:"column:coordinator_name"`
	EstablishedDate  string    `json:"established_date" gorm:"column:established_date"`
	NoOfPreachers    int       `json:"no_of_preachers" gorm:"column:no_of_preachers"`
	Country          string    `json:"country" gorm:"column:country"`
	State            string    `json:"state" gorm:"column:state"`
	City             string    `json:"city" gorm:"column:city"`
	PINCode          string    `json:"pin_code" gorm:"column:pin_code"`
	Address          string    `json:"address" gorm:"column:address"`
	EntryTimestamp   string    `json:"entry_timestamp" gorm:"column:entry_timestamp"`
	UpdatedTimestamp string    `json:"updated_timestamp" gorm:"column:updated_timestamp"`
	ContactNumber    string    `json:"contact_number" gorm:"column:contact_number"`
	Email            string    `json:"email" gorm:"column:email;uniqueIndex"`
	AashramArea      string    `json:"aashram_area" gorm:"column:aashram_area"`
	OpenDays         int       `json:"open_days" gorm:"column:open_days"`
	OpeningTime      string    `json:"opening_time" gorm:"column:opening_time"`
	ClosingTime      string    `json:"closing_time" gorm:"column:closing_time"`
	CreatedOn        string    `json:"created_on" gorm:"column:created_on"`
	UpdatedOn        string    `json:"updated_on" gorm:"column:updated_on"`
	CreatedBy        uuid.UUID `json:"created_by" gorm:"column:created_by"`
	UpdatedBy        uuid.UUID `json:"updated_by" gorm:"column:updated_by"`
}

func (Branch) TableName() string { return "branch" }
