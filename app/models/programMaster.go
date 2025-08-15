package models

import (
	"time"

	"github.com/google/uuid"
)

type ProgramMaster struct {
	ID          uuid.UUID `json:"id"`
	ProgramType string    `json:"program_type"`
	ProgramName string    `json:"program_name"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ProgramMaster) TableName() string { return "ProgramMaster" }
