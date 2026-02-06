package configs

import (
	"github.com/Toppira-Official/backend/internal/domain/repositories"
	"gorm.io/gorm"
)

func NewQuery(db *gorm.DB) *repositories.Query {
	q := repositories.Use(db)
	return q
}
