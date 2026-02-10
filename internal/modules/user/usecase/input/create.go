package input

type CreateUserInput struct {
	Email string
	Phone *string

	Name           *string
	ProfilePicture *string

	Password *string
}
