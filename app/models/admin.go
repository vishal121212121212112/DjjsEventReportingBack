package models

import "github.com/google/uuid"

type Admin struct {
	ID           uuid.UUID `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primary_key"`
	Username     string    `json:"username" gorm:"column:username;unique;not null"`
	Password     string    `json:"password" gorm:"column:password;unique;not null"`
	ContactNumber string   `json:"contact_number" gorm:"column:contact_number"`
	Email        string    `json:"email" gorm:"column:email;unique;not null"`
	CreatedOn    string    `json:"created_on"`
	UpdatedOn    string    `json:"updated_on"`
	CreatedBy    uuid.UUID `json:"created_by"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
}

func (Admin) TableName() string { return "admin_login" }
