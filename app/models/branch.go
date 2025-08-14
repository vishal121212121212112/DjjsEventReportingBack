package models

import "github.com/google/uuid"

type Branch struct {
	ID               uuid.UUID `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primary_key"`
	BranchName       string    `json:"branch_name" gorm:"column:branch_name;not null"`
	ParentBranch     uuid.UUID `json:"parent_branch" gorm:"column:parent_branch"`
	CoordinatorName  string    `json:"coordinator_name" gorm:"column:coordinator_name"`
	EstablishedDate  string    `json:"established_date" gorm:"column:established_date"`
	NoOfPreachers    int       `json:"no_of_preachers" gorm:"column:no_of_preachers"`
	Country          string    `json:"country" gorm:"column:country"`
	State            string    `json:"state" gorm:"column:state"`
	City             string    `json:"city" gorm:"column:city"`
	PinCode          string    `json:"pin_code" gorm:"column:pin_code"`
	Address          string    `json:"address" gorm:"column:address"`
	ContactNumber    string    `json:"contact_number" gorm:"column:contact_number"`
	Email            string    `json:"email" gorm:"column:email;unique"`
	OpenDays         string    `json:"open_days" gorm:"column:open_days"`
	OpeningTime      string    `json:"opening_time" gorm:"column:opening_time"`
	ClosingTime      string    `json:"closing_time" gorm:"column:closing_time"`
	CreatedOn        string    `json:"created_on"`
	UpdatedOn        string    `json:"updated_on"`
	CreatedBy        uuid.UUID `json:"created_by"`
	UpdatedBy        uuid.UUID `json:"updated_by"`
}

func (Branch) TableName() string { return "branch" }
