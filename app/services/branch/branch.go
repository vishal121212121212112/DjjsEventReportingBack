package branchServiceHandler

import (
	database "event-reporting/app/database/pgsql/repository"
	"event-reporting/app/dtos"
	"event-reporting/app/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BranchService struct {
	repo *database.Repository
}

func NewBranchService(repo *database.Repository) *BranchService {
	return &BranchService{repo: repo}
}

// SearchBranches executes the get_branch PostgreSQL function
func (s *BranchService) SearchBranches(req models.BranchSearchRequest) ([]models.BranchSearchResponse, error) {
	var results []models.BranchSearchResponse

	// Convert request to function parameters
	args := s.convertRequestToArgs(req)

	// Execute the PostgreSQL function
	err := s.repo.ExecuteRawFunction("get_branch", args, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to execute get_branch function: %w", err)
	}

	return results, nil
}

// convertRequestToArgs converts the request struct to function parameters
func (s *BranchService) convertRequestToArgs(req models.BranchSearchRequest) []interface{} {
	args := make([]interface{}, 0, 25)

	// Add parameters in the order expected by the function
	args = append(args, req.BranchID)
	args = append(args, req.BranchName)
	args = append(args, req.StateID)
	args = append(args, req.DistrictID)
	args = append(args, req.CountryID)
	args = append(args, req.City)
	args = append(args, req.EstablishedFrom)
	args = append(args, req.EstablishedTo)
	args = append(args, req.CreatedFrom)
	args = append(args, req.CreatedTo)
	args = append(args, req.UpdatedFrom)
	args = append(args, req.UpdatedTo)
	args = append(args, req.InfrastructureType)
	args = append(args, req.SewadarRole)
	args = append(args, req.SewadarGender)
	args = append(args, req.MinTotalSewadars)
	args = append(args, req.MaxTotalSewadars)
	args = append(args, req.MinTotalInfrastructure)
	args = append(args, req.MaxTotalInfrastructure)

	// Set default values for sorting and pagination
	sortBy := "branch_branch_id"
	if req.SortBy != nil {
		sortBy = *req.SortBy
	}
	args = append(args, sortBy)

	sortDir := "ASC"
	if req.SortDir != nil {
		sortDir = *req.SortDir
	}
	args = append(args, sortDir)

	limit := 50
	if req.Limit != nil {
		limit = *req.Limit
	}
	args = append(args, limit)

	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}
	args = append(args, offset)

	return args
}

func (s *BranchService) UpdateBranchDetails(branchID uuid.UUID, req dtos.UpdateBranchDetailsRequest) error {
	// Start a transaction for data consistency
	tx := s.repo.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Update main branch record
	if err := s.updateBranchInfo(tx, branchID, req); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update main branch: %w", err)
	}

	// 2. Create branch districts
	if err := s.createBranchDistricts(tx, branchID, req.Districts, req.UpdatedBy); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create branch districts: %w", err)
	}

	// 3. Create branch areas
	if err := s.createBranchAreas(tx, branchID, req.Areas, req.UpdatedBy); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create branch areas: %w", err)
	}

	// 4. Create branch infrastructures
	if err := s.createBranchInfrastructures(tx, branchID, req.Infrastructures, req.UpdatedBy); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create branch infrastructures: %w", err)
	}

	// 5. Create branch sewadars
	if err := s.createBranchSewadars(tx, branchID, req.Sewadars, req.UpdatedBy); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create branch sewadars: %w", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// updateMainBranch updates the main branch record
func (s *BranchService) updateBranchInfo(tx *gorm.DB, branchID uuid.UUID, req dtos.UpdateBranchDetailsRequest) error {
	updates := map[string]interface{}{
		"branch_name":      req.BranchName,
		"parent_branch":    req.ParentBranch,
		"coordinator_name": req.CoordinatorName,
		"established_date": req.EstablishedDate,
		"no_of_preachers":  req.NoOfPreachers,
		"country":          req.Country,
		"state":            req.State,
		"city":             req.City,
		"pin_code":         req.PINCode,
		"address":          req.Address,
		"contact_number":   req.ContactNumber,
		"aashram_area":     req.AashramArea,
		"open_days":        req.OpenDays,
		"opening_time":     req.OpeningTime,
		"closing_time":     req.ClosingTime,
		"updated_on":       time.Now().UTC(),
		"updated_by":       uuid.MustParse(req.UpdatedBy),
	}

	return tx.Model(&models.Branch{}).Where("id = ?", branchID).Updates(updates).Error
}

// createBranchDistricts creates district records for the branch
func (s *BranchService) createBranchDistricts(tx *gorm.DB, branchID uuid.UUID, districts []dtos.BranchDistrictRequest, updatedBy string) error {
	for _, district := range districts {
		branchDistrict := models.BranchDistrict{
			FkBranchID:       branchID,
			District:         district.District,
			AreaName:         district.AreaName,
			AreaCoverage:     district.AreaCoverage,
			DistrictCoverage: district.DistrictCoverage,
			CreatedOn:        time.Now().UTC(),
			UpdatedOn:        time.Now().UTC(),
			CreatedBy:        uuid.MustParse(updatedBy),
			UpdatedBy:        uuid.MustParse(updatedBy),
		}

		if err := tx.Create(&branchDistrict).Error; err != nil {
			return err
		}
	}
	return nil
}

// createBranchAreas creates area records for the branch
func (s *BranchService) createBranchAreas(tx *gorm.DB, branchID uuid.UUID, areas []dtos.BranchAreaRequest, updatedBy string) error {
	for _, area := range areas {
		branchArea := models.BranchArea{
			FkBranchID:   branchID,
			AreaName:     area.AreaName,
			AreaCoverage: area.AreaCoverage,
			CreatedOn:    time.Now().UTC(),
			UpdatedOn:    time.Now().UTC(),
			CreatedBy:    uuid.MustParse(updatedBy),
			UpdatedBy:    uuid.MustParse(updatedBy),
		}

		if err := tx.Create(&branchArea).Error; err != nil {
			return err
		}
	}
	return nil
}

// createBranchInfrastructures creates infrastructure records for the branch
func (s *BranchService) createBranchInfrastructures(tx *gorm.DB, branchID uuid.UUID, infrastructures []dtos.BranchInfrastructureRequest, updatedBy string) error {
	for _, infra := range infrastructures {
		branchInfra := models.BranchInfrastructure{
			FkBranchID:         branchID,
			InfrastructureType: infra.InfrastructureType,
			Count:              infra.Count,
			CreatedOn:          time.Now().UTC(),
			UpdatedOn:          time.Now().UTC(),
			CreatedBy:          uuid.MustParse(updatedBy),
			UpdatedBy:          uuid.MustParse(updatedBy),
		}

		if err := tx.Create(&branchInfra).Error; err != nil {
			return err
		}
	}
	return nil
}

// createBranchSewadars creates sewadar records for the branch
func (s *BranchService) createBranchSewadars(tx *gorm.DB, branchID uuid.UUID, sewadars []dtos.BranchSamarpitSewadarRequest, updatedBy string) error {
	for _, sewadar := range sewadars {
		branchSewadar := models.BranchSamarpitSewadar{
			FkBranchID:       branchID,
			Role:             sewadar.Role,
			VolunteersName:   sewadar.VolunteersName,
			Gender:           sewadar.Gender,
			Age:              sewadar.Age,
			Responsibility:   sewadar.Responsibility,
			DateOfSamarpit:   sewadar.DateOfSamarpit,
			Qualification:    sewadar.Qualification,
			DOB:              sewadar.DOB,
			ContactNumber:    sewadar.ContactNumber,
			Email:            sewadar.Email,
			DateOfInitiation: sewadar.DateOfInitiation,
			CreatedOn:        time.Now().UTC(),
			UpdatedOn:        time.Now().UTC(),
			CreatedBy:        uuid.MustParse(updatedBy),
			UpdatedBy:        uuid.MustParse(updatedBy),
		}

		if err := tx.Create(&branchSewadar).Error; err != nil {
			return err
		}
	}
	return nil
}

// GetBranchByID retrieves a branch by ID for updating
func (s *BranchService) GetBranchByID(branchID uuid.UUID) (*models.Branch, error) {
	var branch models.Branch
	conditions := map[string]interface{}{"id": branchID}
	if err := s.repo.Find(&branch, conditions); err != nil {
		return nil, fmt.Errorf("branch not found: %w", err)
	}
	return &branch, nil
}

// GetBranchByUserEmail retrieves a branch by user's email
func (s *BranchService) GetBranchByUserEmail(email string) (*models.Branch, error) {
	var branch models.Branch
	conditions := map[string]interface{}{"email": email}
	if err := s.repo.Find(&branch, conditions); err != nil {
		return nil, fmt.Errorf("branch not found for email %s: %w", email, err)
	}
	return &branch, nil
}

func (s *BranchService) CreateBranchWithUserEmail(req dtos.CreateBranchWithUserEmail) error {
	branch := models.Branch{
		Email:     req.Email,
		CreatedOn: time.Now().UTC(),
		UpdatedOn: time.Now().UTC(),
		CreatedBy: uuid.MustParse(req.CreatedBy),
		UpdatedBy: uuid.MustParse(req.UpdatedBy),
	}
	if err := s.repo.Create(&branch); err != nil {
		return fmt.Errorf("failed to create branch: %w", err)
	}

	return nil
}
