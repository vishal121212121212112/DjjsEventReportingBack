package dtos

type CreateBranchWithUserEmail struct {
	Email     string `json:"email" binding:"required,email"`
	CreatedBy string `json:"created_by" binding:"required"`
	UpdatedBy string `json:"updated_by" binding:"required"`
}

// UpdateBranchDetailsRequest represents the complete branch update request
type UpdateBranchDetailsRequest struct {
	// Main Branch Information
	BranchName      string `json:"branch_name"`
	ParentBranch    string `json:"parent_branch"`
	CoordinatorName string `json:"coordinator_name"`
	EstablishedDate string `json:"established_date"`
	NoOfPreachers   int    `json:"no_of_preachers"`

	// Location Information
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	PINCode string `json:"pin_code"`
	Address string `json:"address"`

	// Contact and Operational Info
	ContactNumber string `json:"contact_number"`
	AashramArea   string `json:"aashram_area"`
	OpenDays      int    `json:"open_days"`
	OpeningTime   string `json:"opening_time"`
	ClosingTime   string `json:"closing_time"`

	// Related Entities
	Districts       []BranchDistrictRequest        `json:"districts"`
	Areas           []BranchAreaRequest            `json:"areas"`
	Infrastructures []BranchInfrastructureRequest  `json:"infrastructures"`
	Sewadars        []BranchSamarpitSewadarRequest `json:"sewadars"`

	// Audit
	UpdatedBy string `json:"updated_by" binding:"required"`
}

type BranchDistrictRequest struct {
	District         string `json:"district" binding:"required"`
	AreaName         string `json:"area_name"`
	AreaCoverage     string `json:"area_coverage"`
	DistrictCoverage string `json:"district_coverage"`
}

type BranchAreaRequest struct {
	AreaName     string `json:"area_name" binding:"required"`
	AreaCoverage string `json:"area_coverage"`
}

type BranchInfrastructureRequest struct {
	InfrastructureType string `json:"infrastructure_type" binding:"required"`
	Count              int    `json:"count"`
}

type BranchSamarpitSewadarRequest struct {
	Role             string `json:"role" binding:"required"`
	VolunteersName   string `json:"volunteers_name" binding:"required"`
	Gender           string `json:"gender"`
	Age              int    `json:"age"`
	Responsibility   string `json:"responsibility"`
	DateOfSamarpit   string `json:"date_of_samarpit"`
	Qualification    string `json:"qualification"`
	DOB              string `json:"dob"`
	ContactNumber    string `json:"contact_number"`
	Email            string `json:"email"`
	DateOfInitiation string `json:"date_of_initiation"`
}
