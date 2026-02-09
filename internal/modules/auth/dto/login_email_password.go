package dto

type LoginWithEmailPasswordInput struct {
	Email    string `binding:"required,email" json:"email" example:"user@example.com"`
	Password string `binding:"required,min=8,max=72" json:"password" example:"StrongPassword1234"`
} //	@name	LoginWithEmailPasswordInput
