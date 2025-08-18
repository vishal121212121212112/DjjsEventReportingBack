package geographyServiceHandler

import (
	database "event-reporting/app/database/pgsql/repository"
	"event-reporting/app/dtos"
	"event-reporting/app/models"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type GeographyService struct {
	repo *database.Repository
}

func NewGeographyService(repo *database.Repository) *GeographyService {
	return &GeographyService{repo: repo}
}

func (s *GeographyService) SearchCountries(q string, limit, offset int) (*dtos.PagedCountries, error) {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	q = strings.TrimSpace(q)

	conditions := map[string]interface{}{}
	var raw string
	var args []interface{}

	if q != "" {
		raw = `(name ILIKE ? OR iso2 ILIKE ? OR iso3 ILIKE ? OR phone_code ILIKE ?)`
		arg := "%" + q + "%"
		args = []interface{}{arg, arg, arg, arg}
	}

	total, err := s.repo.CountWithRawCondition(&models.Country{}, conditions, raw, args)
	if err != nil {
		return nil, fmt.Errorf("count countries: %w", err)
	}

	var items []models.Country
	if err := s.repo.FindAllWithRawConditionAndOrder(
		&items, conditions, raw, args, "name ASC", limit, offset,
	); err != nil {
		return nil, fmt.Errorf("list countries: %w", err)
	}

	return &dtos.PagedCountries{
		Items:  dtos.ToCountryDTOs(items),
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (s *GeographyService) SearchStates(countryID string, q string, limit, offset int) (*dtos.PagedStates, error) {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	conditions := map[string]interface{}{}
	if cid := strings.TrimSpace(countryID); cid != "" {
		parsed, err := uuid.Parse(cid)
		if err != nil {
			return nil, fmt.Errorf("invalid country_id: %w", err)
		}
		conditions["country_id"] = parsed
	}

	var raw string
	var args []interface{}
	if q = strings.TrimSpace(q); q != "" {
		raw = `(name ILIKE ? OR code ILIKE ?)`
		arg := "%" + q + "%"
		args = []interface{}{arg, arg}
	}

	total, err := s.repo.CountWithRawCondition(&models.State{}, conditions, raw, args)
	if err != nil {
		return nil, fmt.Errorf("count states: %w", err)
	}

	var items []models.State
	if err := s.repo.FindAllWithRawConditionAndOrder(
		&items, conditions, raw, args, "name ASC", limit, offset,
	); err != nil {
		return nil, fmt.Errorf("list states: %w", err)
	}

	return &dtos.PagedStates{
		Items:  dtos.ToStateDTOs(items),
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (s *GeographyService) SearchDistricts(stateID string, q string, limit, offset int) (*dtos.PagedDistricts, error) {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	conditions := map[string]interface{}{}
	if sid := strings.TrimSpace(stateID); sid != "" {
		parsed, err := uuid.Parse(sid)
		if err != nil {
			return nil, fmt.Errorf("invalid state_id: %w", err)
		}
		conditions["state_id"] = parsed
	}

	var raw string
	var args []interface{}
	if q = strings.TrimSpace(q); q != "" {
		raw = `name ILIKE ?`
		args = []interface{}{"%" + q + "%"}
	}

	total, err := s.repo.CountWithRawCondition(&models.District{}, conditions, raw, args)
	if err != nil {
		return nil, fmt.Errorf("count districts: %w", err)
	}

	var items []models.District
	if err := s.repo.FindAllWithRawConditionAndOrder(
		&items, conditions, raw, args, "name ASC", limit, offset,
	); err != nil {
		return nil, fmt.Errorf("list districts: %w", err)
	}

	return &dtos.PagedDistricts{
		Items:  dtos.ToDistrictDTOs(items),
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (s *GeographyService) SearchCities(districtID string, q string, limit, offset int) (*dtos.PagedCities, error) {
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}

	conditions := map[string]interface{}{}
	if did := strings.TrimSpace(districtID); did != "" {
		parsed, err := uuid.Parse(did)
		if err != nil {
			return nil, fmt.Errorf("invalid district_id: %w", err)
		}
		conditions["district_id"] = parsed
	}

	var raw string
	var args []interface{}
	if q = strings.TrimSpace(q); q != "" {
		raw = `(name ILIKE ? OR pincode ILIKE ?)`
		arg := "%" + q + "%"
		args = []interface{}{arg, arg}
	}

	total, err := s.repo.CountWithRawCondition(&models.City{}, conditions, raw, args)
	if err != nil {
		return nil, fmt.Errorf("count cities: %w", err)
	}

	var items []models.City
	if err := s.repo.FindAllWithRawConditionAndOrder(
		&items, conditions, raw, args, "name ASC", limit, offset,
	); err != nil {
		return nil, fmt.Errorf("list cities: %w", err)
	}

	return &dtos.PagedCities{
		Items:  dtos.ToCityDTOs(items),
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}, nil
}
