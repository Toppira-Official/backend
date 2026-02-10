package dto

type UpdateMeInput struct {
	Name     *string `json:"name,omitempty" example:"John Doe"`
	Phone    *string `json:"phone,omitempty" example:"09123456789"`
	Password *string `binding:"omitempty,min=8,max=72" json:"password,omitempty" example:"securepassword"`
} //	@name	UpdateMeInput
