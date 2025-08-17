package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
Hierarchy:
Country 1—N State 1—N District 1—N City
(You can also let City reference State directly if you need, but District → City is common in India.)
*/

// ---------- Country ----------
type Country struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(120);not null;uniqueIndex:ux_country_name"`
	ISO2      string    `gorm:"type:char(2);not null;uniqueIndex:ux_country_iso2"`
	ISO3      string    `gorm:"type:char(3);not null;uniqueIndex:ux_country_iso3"`
	PhoneCode string    `gorm:"type:varchar(8)"`
	Currency  string    `gorm:"type:varchar(8)"`
	States    []State   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // <-- CASCADE
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Country) TableName() string { return "countries" }

// ---------- State ----------
type State struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CountryID uuid.UUID `gorm:"type:uuid;not null;index"`
	// Explicit FK with cascade
	Country Country `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CountryID;references:ID"`

	Name string `gorm:"type:varchar(120);not null;uniqueIndex:ux_state_name_country,priority:2"`
	Code string `gorm:"type:varchar(10);not null;uniqueIndex:ux_state_code_country,priority:2"`
	// composite uniqueness scopes
	// ux_state_name_country -> (country_id, name)
	// ux_state_code_country -> (country_id, code)
	// add the parent part:
	// (GORM uses the same index name with priority to compose)

	Districts []District `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // <-- CASCADE
	Cities    []City     `gorm:"-"`                                            // exposed via queries if needed
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (State) TableName() string { return "states" }

// ---------- District ----------
type District struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	StateID uuid.UUID `gorm:"type:uuid;not null;index"`
	State   State     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:StateID;references:ID"`

	Name string `gorm:"type:varchar(150);not null;uniqueIndex:ux_district_name_state,priority:2"`
	// ux_district_name_state -> (state_id, name)

	Cities    []City `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // <-- CASCADE
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (District) TableName() string { return "districts" }

// ---------- City ----------
type City struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	DistrictID uuid.UUID `gorm:"type:uuid;not null;index"`
	District   District  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DistrictID;references:ID"`

	Name string `gorm:"type:varchar(150);not null;uniqueIndex:ux_city_name_district,priority:2"`
	// ux_city_name_district -> (district_id, name)

	Pincode   string   `gorm:"type:varchar(10);index"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	IsCapital bool     `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (City) TableName() string { return "cities" }
