package handler

import (
	"net/http"

	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/notification/handler/dto/input"
	usecaseInput "github.com/Toppira-Official/Reminder_Server/internal/modules/notification/usecase/input"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/notification/usecase"

	sharedDto "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type SubscribeFirebaseHandler struct {
	subscribeFirebaseUsecase usecase.SubscribeFirebaseUsecase
	envs                     configs.Environments
}

func NewSubscribeFirebaseHandler(
	subscribeFirebaseUsecase usecase.SubscribeFirebaseUsecase,
	envs configs.Environments,
) *SubscribeFirebaseHandler {
	return &SubscribeFirebaseHandler{
		subscribeFirebaseUsecase: subscribeFirebaseUsecase,
		envs:                     envs,
	}
}

// Subscribe godoc
//
//	@Summary		Subscribe to Firebase Cloud Messaging
//	@Description	Subscribe to Firebase Cloud Messaging with the provided token
//	@Accept			json
//	@Produce		json
//	@Tags			Notification
//	@Param			body	body		input.SubscribeFirebaseInput	true	"Subscribe Firebase Input"
//	@Success		201		{object}	sharedDto.HttpOutput[any]
//	@Failure		401		{object}	apperrors.ClientError
//	@Failure		500		{object}	apperrors.ClientError
//	@Failure		503		{object}	apperrors.ClientError
//	@Router			/notification/firebase/subscribe [post]
func (h *SubscribeFirebaseHandler) Subscribe(c *gin.Context) {
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

	var input input.SubscribeFirebaseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperrors.E(apperrors.ErrUserInvalidData, err))
		return
	}

	usecaseInput := &usecaseInput.SubscribeFirebaseInput{
		UserID: user.ID,
		Token:  input.Token,
	}
	if _, err := h.subscribeFirebaseUsecase.Execute(ctx, usecaseInput); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, sharedDto.HttpOutput[any]{})
}
