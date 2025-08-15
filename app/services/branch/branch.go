package branchServiceHandler

import (
	database "event-reporting/app/database/pgsql/repository"
	"event-reporting/app/models"
	"fmt"
)

type BranchService struct {
	repo *database.Repository
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
