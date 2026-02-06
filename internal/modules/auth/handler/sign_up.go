package handler

import (
	"net/http"

	userUsecase "github.com/Toppira-Official/backend/internal/modules/user/usecase"
	"github.com/gin-gonic/gin"
)

type SignUpHandler struct {
	createUserUsecase userUsecase.CreateUserUsecase
}

func NewSignUpHandler(createUserUsecase userUsecase.CreateUserUsecase) *SignUpHandler {
	return &SignUpHandler{createUserUsecase: createUserUsecase}
}

// SignUpWithEmailPassword godoc
//
//	@Summary	sign up with email and password
//	@Tags		Authentication
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	SignUpWithEmailPasswordOutput
//	@Failure	400	{object}	SignUpWithEmailPasswordOutput
//	@Failure	500	{object}	SignUpWithEmailPasswordOutput
//	@Router		/auth/sign-up-with-user-password [post]
func (hl *SignUpHandler) SignUpWithEmailPassword(c *gin.Context) {
	var input SignUpWithEmailPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, SignUpWithEmailPasswordOutput{Message: "invalid body request"})
		return
	}

	user, err := hl.createUserUsecase.Execute(c.Request.Context(), input.MapUser())
	if err != nil {
		c.JSON(http.StatusInternalServerError, SignUpWithEmailPasswordOutput{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SignUpWithEmailPasswordOutput{Message: "welcome",
		Data: map[string]any{
			"user": user,
		}})
}
