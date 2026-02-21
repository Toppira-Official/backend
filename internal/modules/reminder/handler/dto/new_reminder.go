package dto

import (
	"time"

	"github.com/Toppira-Official/Reminder_Server/internal/shared/constants"
)

type NewReminderInput struct {
	Title         string                      `json:"title" binding:"required,min=3,max=200"`
	Description   *string                     `json:"description,omitempty"`
	ReminderTimes []time.Time                 `json:"reminder_times,omitempty"`
	ScheduledAt   time.Time                   `json:"scheduled_at" binding:"required"`
	Priority      *constants.ReminderPriority `json:"priority,omitempty"`
} //	@name	NewReminderInput
