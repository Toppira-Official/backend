package configs

import (
	"github.com/Toppira-Official/Reminder_Server/internal/shared/repositories"
	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *repositories.Query {
	q := repositories.Use(db)
	return q
}
