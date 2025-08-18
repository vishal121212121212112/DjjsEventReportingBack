package dtos

import (
	"event-reporting/app/models"
	"time"

	"github.com/google/uuid"
)

type CountryDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ISO2      string    `json:"iso2"`
	ISO3      string    `json:"iso3"`
	PhoneCode string    `json:"phone_code,omitempty"`
	Currency  string    `json:"currency,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StateDTO struct {
	ID        uuid.UUID `json:"id"`
	CountryID uuid.UUID `json:"country_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DistrictDTO struct {
	ID        uuid.UUID `json:"id"`
	StateID   uuid.UUID `json:"state_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CityDTO struct {
	ID         uuid.UUID `json:"id"`
	DistrictID uuid.UUID `json:"district_id"`
	Name       string    `json:"name"`
	Pincode    string    `json:"pincode,omitempty"`
	Latitude   *float64  `json:"latitude,omitempty"`
	Longitude  *float64  `json:"longitude,omitempty"`
	IsCapital  bool      `json:"is_capital"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PagedCountries struct {
	Items  []CountryDTO `json:"items"`
	Total  int64        `json:"total"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
}

type PagedStates struct {
	Items  []StateDTO `json:"items"`
	Total  int64      `json:"total"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
}

type PagedDistricts struct {
	Items  []DistrictDTO `json:"items"`
	Total  int64         `json:"total"`
	Limit  int           `json:"limit"`
	Offset int           `json:"offset"`
}

type PagedCities struct {
	Items  []CityDTO `json:"items"`
	Total  int64     `json:"total"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}

func ToCountryDTO(m models.Country) CountryDTO {
	return CountryDTO{
		ID:        m.ID,
		Name:      m.Name,
		ISO2:      m.ISO2,
		ISO3:      m.ISO3,
		PhoneCode: m.PhoneCode,
		Currency:  m.Currency,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToStateDTO(m models.State) StateDTO {
	return StateDTO{
		ID:        m.ID,
		CountryID: m.CountryID,
		Name:      m.Name,
		Code:      m.Code,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToDistrictDTO(m models.District) DistrictDTO {
	return DistrictDTO{
		ID:        m.ID,
		StateID:   m.StateID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToCityDTO(m models.City) CityDTO {
	return CityDTO{
		ID:         m.ID,
		DistrictID: m.DistrictID,
		Name:       m.Name,
		Pincode:    m.Pincode,
		Latitude:   m.Latitude,
		Longitude:  m.Longitude,
		IsCapital:  m.IsCapital,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

type CommonQuery struct {
	Q      string `form:"q"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type StatesQuery struct {
	CommonQuery
	CountryID string `form:"country_id"`
}

type DistrictsQuery struct {
	CommonQuery
	StateID string `form:"state_id"`
}

type CitiesQuery struct {
	CommonQuery
	DistrictID string `form:"district_id"`
}

// Slice helpers
func ToCountryDTOs(ms []models.Country) []CountryDTO {
	out := make([]CountryDTO, len(ms))
	for i, m := range ms {
		out[i] = ToCountryDTO(m)
	}
	return out
}

func ToStateDTOs(ms []models.State) []StateDTO {
	out := make([]StateDTO, len(ms))
	for i, m := range ms {
		out[i] = ToStateDTO(m)
	}
	return out
}

func ToDistrictDTOs(ms []models.District) []DistrictDTO {
	out := make([]DistrictDTO, len(ms))
	for i, m := range ms {
		out[i] = ToDistrictDTO(m)
	}
	return out
}

func ToCityDTOs(ms []models.City) []CityDTO {
	out := make([]CityDTO, len(ms))
	for i, m := range ms {
		out[i] = ToCityDTO(m)
	}
	return out
}
