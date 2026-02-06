package entities

import (
	"time"

	"github.com/Toppira-Official/backend/internal/domain/constants"
	"gorm.io/datatypes"
)

type Reminder struct {
	Base `gorm:"embedded"`

	Title       string  `gorm:"not null" json:"title"`
	Description *string `json:"description,omitempty"`

	Status constants.ReminderStatus `gorm:"type:varchar(20);not null" json:"status"`

	ReminderTimes datatypes.JSON `gorm:"type:json" json:"reminder_times,omitempty"`

	ScheduledAt time.Time `gorm:"not null;index" json:"scheduled_at"`

	Priority *string `gorm:"type:varchar(20)" json:"priority"`

	UserID uint `gorm:"not null;index" json:"user_id"`
	User   User `gorm:"constraint:OnDelete:CASCADE"`
}
