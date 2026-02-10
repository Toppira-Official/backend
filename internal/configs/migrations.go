package configs

import (
	"github.com/Toppira-Official/backend/internal/shared/entities"
	"gorm.io/gorm"
)

func LoadMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.User{},
		&entities.Reminder{},
	)
}
