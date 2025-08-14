package eventHistoryService

import (
	"context"
	"errors"
	"event-reporting/app/database/pgsql/connection"
	database "event-reporting/app/database/pgsql/repository"
	"event-reporting/app/models"
	"fmt"
	"time"
)

type EventHistoryService struct {
	repo *database.Repository
}

// CreateEventHistory inserts a new EventHistory record into the database
func (s *EventHistoryService) CreateEventHistory(ctx context.Context, req *models.EventHistory) (*models.EventHistory, error) {
	// Set creation metadata
	req.CreatedOn = time.Now().UTC()
	req.UpdatedOn = time.Now().UTC()

	// ------ Event History
	var existing models.EventHistory
	if err := connection.Db.
		Where("fk_program_master_id = ? AND fk_branch_id = ? AND start_date = ? AND is_deleted = false",
			req.FKProgramMasterID, req.FKBranchID, req.StartDate).
		First(&existing).Error; err == nil {
		return nil, errors.New("event history already exists for this program, branch, and start date")
	}

	// Insert the new record
	if err := connection.Db.Create(req).Error; err != nil {
		return nil, err
	}

	// Retrieve the newly created record (optional, but ensures DB defaults are loaded)
	var created models.EventHistory
	if err := connection.Db.First(&created, req.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch created event history: %w", err)
	}

	return &created, nil
}
