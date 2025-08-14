package models

import "time"

type MediaAndDocumentation struct {
    ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
    FKProgramID   int       `gorm:"not null" json:"fk_program_id"`
    FileUrl       string    `gorm:"type:varchar(255);null" json:"file_url"`
    FileName      string    `gorm:"type:varchar(255);not null" json:"file_name"`
    FileSize      int       `gorm:"null" json:"file_size"`
    ContentType   string    `gorm:"type:varchar(50);null" json:"content_type"`
    CreatedOn     time.Time `gorm:"default:current_timestamp" json:"created_on"`
    UpdatedOn     time.Time `gorm:"default:current_timestamp" json:"updated_on"`
    CreatedBy     string    `gorm:"type:varchar(255);null" json:"created_by"`
    UpdatedBy     string    `gorm:"type:varchar(255);null" json:"updated_by"`
}

func (MediaAndDocumentation) TableName() string { return "MediaAndDocumentation" }

