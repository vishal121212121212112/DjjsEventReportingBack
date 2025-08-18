package models

// type Admin struct {
// 	ID            uuid.UUID `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primary_key"`
// 	Username      string    `json:"username" gorm:"column:username;unique;not null"`
// 	Password      string    `json:"password" gorm:"column:password;unique;not null"`
// 	ContactNumber string    `json:"contact_number" gorm:"column:contact_number"`
// 	Email         string    `json:"email" gorm:"column:email;unique;not null"`
// 	CreatedOn     string    `json:"created_on"`
// 	UpdatedOn     string    `json:"updated_on"`
// 	CreatedBy     uuid.UUID `json:"created_by"`
// 	UpdatedBy     uuid.UUID `json:"updated_by"`
// }

// func (Admin) TableName() string { return "admin_login" }

// type User struct {
// 	ID            uuid.UUID `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid();primary_key"`
// 	Username      string    `json:"username" gorm:"column:username;unique"` // For Admins
// 	Email         string    `json:"email" gorm:"column:email;unique"`       // For Admins & Branches
// 	Password      string    `json:"password" gorm:"column:password;not null"`
// 	ContactNumber string    `json:"contact_number" gorm:"column:contact_number"`
// 	Type          string    `json:"type" gorm:"column:type"` // Admin/Branch
// 	Token         string    `json:"token" gorm:"column:token"`
// 	ExpiredOn     string    `json:"expired_on" gorm:"column:expired_on"`
// 	LastLoginOn   string    `json:"last_login_on" gorm:"column:last_login_on"`
// 	FirstLoginOn  string    `json:"first_login_on" gorm:"column:first_login_on"`
// 	CreatedOn     string    `json:"created_on" gorm:"column:created_on"`
// 	UpdatedOn     string    `json:"updated_on" gorm:"column:updated_on"`
// 	CreatedBy     uuid.UUID `json:"created_by" gorm:"column:created_by"`
// 	UpdatedBy     uuid.UUID `json:"updated_by" gorm:"column:updated_by"`
// }

// func (User) TableName() string { return "user_login" }
