package handler

import (
	"net/http"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/handler/dto"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/reminder/usecase"
	output "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type MyRemindersHandler struct {
	listRemindersUsecase usecase.ListRemindersUsecase
}

func NewMyRemindersHandler(listRemindersUsecase usecase.ListRemindersUsecase) *MyRemindersHandler {
	return &MyRemindersHandler{listRemindersUsecase: listRemindersUsecase}
}

// MyReminders godoc
//
//	@Summary	returns user's all reminders
//	@Tags		Reminder
//	@Produce	json
//	@Param		page	query		int	false	"Page number"		default(1)
//	@Param		limit	query		int	false	"Items per page"	default(10)
//	@Success	200		{object}	output.HttpOutput
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Failure	503		{object}	apperrors.ClientError
//	@Security	BearerAuth
//	@Router		/reminder [get]
func (hl *MyRemindersHandler) MyReminders(c *gin.Context) {
	var q output.PaginationInput
	if err := c.ShouldBindQuery(&q); err != nil {
		c.Error(apperrors.E(apperrors.ErrReminderInvalidData))
		return
	}

	if q.Page == 0 {
		q.Page = 1
	}
	if q.Limit == 0 {
		q.Limit = 10
	}

	ctx := c.Request.Context()

	userVal, exists := c.Get("user")
	if !exists {
		c.Error(apperrors.E(apperrors.ErrUserNotFound))
		return
	}

	user, ok := userVal.(*entities.User)
	if !ok {
		c.Error(apperrors.E(apperrors.ErrUserNotFound))
		return
	}

	reminders, err := hl.listRemindersUsecase.Execute(ctx, user.ID, q.Page, q.Limit)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output.HttpOutput[dto.MyRemindersOutput]{
		Data: dto.MyRemindersOutput{
			Reminders: reminders,
		},
	})
}
