package input

import (
	"time"

	"github.com/Toppira-Official/Reminder_Server/internal/shared/constants"
)

type CreateReminderInput struct {
	Title         string
	Description   *string
	ReminderTimes []time.Time
	ScheduledAt   time.Time
	UserID        uint
	Priority      *constants.ReminderPriority
}
