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

// BranchSearchResponse represents the response from get_branch function
type BranchSearchResponse struct {
	BranchBranchID                   *int64  `json:"branch_branch_id" gorm:"column:branch_branch_id"`
	BranchBranchName                 *string `json:"branch_branch_name" gorm:"column:branch_branch_name"`
	BranchStateID                    *int64  `json:"branch_state_id" gorm:"column:branch_state_id"`
	BranchDistrictID                 *int64  `json:"branch_district_id" gorm:"column:branch_district_id"`
	BranchCountryID                  *int64  `json:"branch_country_id" gorm:"column:branch_country_id"`
	BranchCity                       *string `json:"branch_city" gorm:"column:branch_city"`
	BranchEstablishedDate            *string `json:"branch_established_date" gorm:"column:branch_established_date"`
	BranchCreatedOn                  *string `json:"branch_created_on" gorm:"column:branch_created_on"`
	BranchUpdatedOn                  *string `json:"branch_updated_on" gorm:"column:branch_updated_on"`
	StateStateName                   *string `json:"state_state_name" gorm:"column:state_state_name"`
	DistrictDistrictName             *string `json:"district_district_name" gorm:"column:district_district_name"`
	InfrastructureInfrastructureType *string `json:"infrastructure_infrastructure_type" gorm:"column:infrastructure_infrastructure_type"`
	SewaSewadarName                  *string `json:"sewa_sewadar_name" gorm:"column:sewa_sewadar_name"`
	SewaSewadarRole                  *string `json:"sewa_sewadar_role" gorm:"column:sewa_sewadar_role"`
	SewaSewadarGender                *string `json:"sewa_sewadar_gender" gorm:"column:sewa_sewadar_gender"`
	TotalSewadarCount                *int64  `json:"total_sewadar_count" gorm:"column:total_sewadar_count"`
	TotalInfrastructureCount         *int64  `json:"total_infrastructure_count" gorm:"column:total_infrastructure_count"`
}

// BranchSearchRequest represents the parameters for get_branch function
type BranchSearchRequest struct {
	BranchID               *int64  `json:"branch_id,omitempty"`
	BranchName             *string `json:"branch_name,omitempty"`
	StateID                *int64  `json:"state_id,omitempty"`
	DistrictID             *int64  `json:"district_id,omitempty"`
	CountryID              *int64  `json:"country_id,omitempty"`
	City                   *string `json:"city,omitempty"`
	EstablishedFrom        *string `json:"established_from,omitempty"`
	EstablishedTo          *string `json:"established_to,omitempty"`
	CreatedFrom            *string `json:"created_from,omitempty"`
	CreatedTo              *string `json:"created_to,omitempty"`
	UpdatedFrom            *string `json:"updated_from,omitempty"`
	UpdatedTo              *string `json:"updated_to,omitempty"`
	InfrastructureType     *string `json:"infrastructure_type,omitempty"`
	SewadarRole            *string `json:"sewadar_role,omitempty"`
	SewadarGender          *string `json:"sewadar_gender,omitempty"`
	MinTotalSewadars       *int    `json:"min_total_sewadars,omitempty"`
	MaxTotalSewadars       *int    `json:"max_total_sewadars,omitempty"`
	MinTotalInfrastructure *int    `json:"min_total_infrastructure,omitempty"`
	MaxTotalInfrastructure *int    `json:"max_total_infrastructure,omitempty"`
	SortBy                 *string `json:"sort_by,omitempty"`
	SortDir                *string `json:"sort_dir,omitempty"`
	Limit                  *int    `json:"limit,omitempty"`
	Offset                 *int    `json:"offset,omitempty"`
}
