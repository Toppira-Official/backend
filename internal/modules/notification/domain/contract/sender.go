package contract

import (
	"context"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/notification/domain/model"
)

type Sender interface {
	Send(ctx context.Context, message model.Message) error
}
