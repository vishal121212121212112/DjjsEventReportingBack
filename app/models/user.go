package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey"`
	Username      string    `json:"username" gorm:"column:username;uniqueIndex"`   // For Admins
	Email         string    `json:"email" gorm:"column:email;uniqueIndex"`         // For Admins & Branches
	FKBranchEmail string    `json:"fk_branch_email" gorm:"column:fk_branch_email"` // For Branch users, can be blank for Admins
	Password      string    `json:"password" gorm:"column:password;not null"`
	ContactNumber string    `json:"contact_number" gorm:"column:contact_number"`
	Type          string    `json:"type" gorm:"column:type"` // Admin/Branch
	Token         string    `json:"token" gorm:"column:token"`
	ExpiredOn     string    `json:"expired_on" gorm:"column:expired_on"`
	LastLoginOn   string    `json:"last_login_on" gorm:"column:last_login_on"`
	FirstLoginOn  string    `json:"first_login_on" gorm:"column:first_login_on"`
	CreatedOn     string    `json:"created_on" gorm:"column:created_on"`
	UpdatedOn     string    `json:"updated_on" gorm:"column:updated_on"`
	CreatedBy     uuid.UUID `json:"created_by" gorm:"column:created_by"`
	UpdatedBy     uuid.UUID `json:"updated_by" gorm:"column:updated_by"`
}

type UserCreateRequest struct {
	Email         string `json:"email" binding:"required,email"`
	Type          string `json:"type" binding:"required"` // Admin, BranchCoordinator, BranchAssistant, ITAssistant
	Username      string `json:"username,omitempty" binding:"required"`
	Password      string `json:"password,omitempty" binding:"required"`
	BranchID      string `json:"branchId,omitempty"`
	ContactNumber string `json:"contactNumber,omitempty"`
}
