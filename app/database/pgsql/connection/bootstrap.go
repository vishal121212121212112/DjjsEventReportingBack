package connection

import (
	"gorm.io/gorm"
	"event-reporting/app/models"
)

func Bootstrap(db *gorm.DB) error {
	// Needed for gen_random_uuid(); safe if already exists
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`).Error; err != nil {
		return err
	}

	// Add ALL your models here
	return db.AutoMigrate(
		&models.User{},
		// &models.Project{}, &models.Contact{},
	)
}
