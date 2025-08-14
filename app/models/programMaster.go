package models

import (
	"time"
)

type ProgramMaster struct {
	ID           int       `json:"id"`           
	ProgramType  string    `json:"program_type"`  
	ProgramName  string    `json:"program_name"`        
	CreatedAt    time.Time `json:"created_at"`   
	UpdatedAt    time.Time `json:"updated_at"`    
}

func (ProgramMaster) TableName() string { return "ProgramMaster" }
