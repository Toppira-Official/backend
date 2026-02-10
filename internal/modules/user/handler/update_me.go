package handler

import (
	"net/http"

	"github.com/Toppira-Official/backend/internal/modules/user/handler/dto"
	"github.com/Toppira-Official/backend/internal/modules/user/usecase"
	userInput "github.com/Toppira-Official/backend/internal/modules/user/usecase/input"
	output "github.com/Toppira-Official/backend/internal/shared/dto"
	"github.com/Toppira-Official/backend/internal/shared/entities"
	apperrors "github.com/Toppira-Official/backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type UpdateMeHandler struct {
	updateUserUsecase usecase.UpdateUserUsecase
}

func NewUpdateMeHandler(updateUserUsecase usecase.UpdateUserUsecase) *UpdateMeHandler {
	return &UpdateMeHandler{
		updateUserUsecase: updateUserUsecase,
	}
}

// UpdateMyInfo godoc
//
//	@Summary	update my(user) data
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		body	body		dto.UpdateMeInput	true	"Update Me Input"
//	@Success	200		{object}	output.HttpOutput
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	401		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Failure	503		{object}	apperrors.ClientError
//	@Security	BearerAuth
//	@Router		/user/me [patch]
func (hl *UpdateMeHandler) UpdateMyInfo(c *gin.Context) {
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

	var input dto.UpdateMeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperrors.E(apperrors.ErrUserInvalidData, err))
		return
	}

	usecaseInput := &userInput.UpdateUserInput{
		ID:       user.ID,
		Name:     input.Name,
		Password: input.Password,
		Phone:    input.Phone,
	}
	updatedUser, err := hl.updateUserUsecase.Execute(ctx, usecaseInput)
	if err != nil {
		c.Error(err)
		return
	}

	updatedUser.Password = nil
	c.JSON(http.StatusOK, output.HttpOutput{
		Data: map[string]any{
			"user": updatedUser,
		},
	})
}
