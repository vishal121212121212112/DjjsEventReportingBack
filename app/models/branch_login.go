package models

import "github.com/google/uuid"

type BranchLogin struct {
	ID            uuid.UUID `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primary_key"`
	FKBranchEmail string    `json:"fk_branch_email" gorm:"column:fk_branch_email"`
	Password      string    `json:"password" gorm:"column:password"`
	Type          string    `json:"type" gorm:"column:type"`
	Token         string    `json:"token" gorm:"column:token"`
	ExpiredOn     string    `json:"expired_on" gorm:"column:expired_on"`
	LastLoginOn   string    `json:"last_login_on" gorm:"column:last_login_on"`
	FirstLoginOn  string    `json:"first_login_on" gorm:"column:first_login_on"`
	CreatedOn     string    `json:"created_on"`
	UpdatedOn     string    `json:"updated_on"`
	CreatedBy     uuid.UUID `json:"created_by"`
	UpdatedBy     uuid.UUID `json:"updated_by"`
}

func (BranchLogin) TableName() string { return "branch_login" }
