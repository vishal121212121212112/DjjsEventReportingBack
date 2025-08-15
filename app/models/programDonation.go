package models

import (
	"time"

	"github.com/google/uuid"
)

type ProgramDonation struct {
	ID                     uuid.UUID `gorm:"primaryKey" json:"id"`
	FKProgramID            string    `gorm:"foreignKey" json:"fk_program_id"`
	FKBranchID             string    `gorm:"foreignKey" json:"fk_branch_id"`
	DonationType           string    `gorm:"type:varchar(255)" json:"donation_type"`
	TotalAmount            float64   `gorm:"type:decimal(10,2)" json:"total_amount"`
	Description            string    `gorm:"type:text" json:"description"`
	EstimatedMaterialValue float64   `gorm:"type:decimal(10,2)" json:"estimated_material_value"`
	CreatedOn              time.Time `gorm:"autoCreateTime" json:"created_on"`
	UpdatedOn              time.Time `gorm:"autoUpdateTime" json:"updated_on"`
	CreatedBy              string    `gorm:"type:varchar(255)" json:"created_by"`
	UpdatedBy              string    `gorm:"type:varchar(255)" json:"updated_by"`
}

func (ProgramDonation) TableName() string { return "ProgramDonation" }
