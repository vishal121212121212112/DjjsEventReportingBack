package models

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `json:"name" gorm:"type:text;not null"`
	Email        string    `json:"email" gorm:"type:text;uniqueIndex;not null"`
	PasswordHash string    `json:"-"    gorm:"type:text;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"not null;default:now()"`
}

func (User) TableName() string { return "users" }
