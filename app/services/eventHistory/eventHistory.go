package EventHistoryService

import (
	"context"
	"event-reporting/app/database/pgsql/connection"
	database "event-reporting/app/database/pgsql/repository"
	"event-reporting/app/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventHistoryService struct {
	repo *database.Repository
}

func (s *EventHistoryService) CreateFullEvent(ctx context.Context, req *models.CreateFullEventRequest) (*models.CreateFullEventRequest, error) {
    return s.withTransaction(ctx, func(tx *gorm.DB) (*models.CreateFullEventRequest, error) {
        // --- Set IDs for EventHistory ---
        req.EventHistory.ID = uuid.New()   
        req.EventHistory.CreatedOn = time.Now().UTC()
        req.EventHistory.UpdatedOn = time.Now().UTC()

        if err := tx.Create(&req.EventHistory).Error; err != nil {
            return nil, fmt.Errorf("failed to create event history: %w", err)
        }
	

        // --- Set IDs for GuestMasters ---
		req.GuestMasters.ID = uuid.New()   
        req.GuestMasters.CreatedOn = time.Now().UTC()
        req.GuestMasters.UpdatedOn = time.Now().UTC()
			
        if err := tx.Create(&req.GuestMasters).Error; err != nil {
            return nil, fmt.Errorf("failed to create event history: %w", err)
        }
		
        // --- Set IDs for MediaAndDocs ---
		req.MediaAndDocs.ID = uuid.New()   
        req.MediaAndDocs.CreatedOn = time.Now().UTC()
        req.MediaAndDocs.UpdatedOn = time.Now().UTC()
			
        if err := tx.Create(&req.MediaAndDocs).Error; err != nil {
            return nil, fmt.Errorf("failed to create MediaAndDocs: %w", err)
        }

        // --- Set IDs for ProgramDonations ---
		req.ProgramDonations.ID = uuid.New()   
        req.ProgramDonations.CreatedOn = time.Now().UTC()
        req.ProgramDonations.UpdatedOn = time.Now().UTC()
			
        if err := tx.Create(&req.ProgramDonations).Error; err != nil {
            return nil, fmt.Errorf("failed to create ProgramDonations: %w", err)
        }

        // --- Set IDs for ProgramVolunteers ---
		req.ProgramVolunteers.ID = uuid.New()   
        req.ProgramVolunteers.CreatedOn = time.Now().UTC()
        req.ProgramVolunteers.UpdatedOn = time.Now().UTC()
			
        if err := tx.Create(&req.ProgramVolunteers).Error; err != nil {
            return nil, fmt.Errorf("failed to create Program Volunteers: %w", err)
        }

        // --- Set IDs for Program Master ---
		req.ProgramMasters.ID = uuid.New()   
        req.ProgramMasters.CreatedOn = time.Now().UTC()
        req.ProgramMasters.UpdatedOn = time.Now().UTC()
			
        if err := tx.Create(&req.ProgramMasters).Error; err != nil {
            return nil, fmt.Errorf("failed to create program master: %w", err)
        }

        // Commit the transaction at the end
        return req, nil
    })
}


//withTransaction
func (s *EventHistoryService) withTransaction(ctx context.Context, fn func(tx *gorm.DB) (*models.CreateFullEventRequest, error)) (*models.CreateFullEventRequest, error) {
	tx := connection.Db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	result, err := fn(tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return result, nil
}

