package types

const (
	UserPasswordMinLength = 8
)

type User struct {
	Email    string
	Password string
}

type AuthenticatedUser struct {
	Email    string
	LoggedIn bool
}

type RegisterUserValues struct {
	Email    string
	Password string
}
type RegisterUserErrors struct {
	Email    string
	Password string
	Other    string
}

type LoginUserValues struct {
	Email    string
	Password string
}
type LoginUserErrors struct {
	Email    string
	Password string
	Other    string
}

func (values RegisterUserValues) Validate() *RegisterUserErrors {
	errors := &RegisterUserErrors{}

	// TODO: validate through validator pkg
	return errors
}

func (errors *RegisterUserErrors) SetOtherError(message string) {
	errors.Other = message
}

func (values LoginUserValues) Validate() *LoginUserErrors {
	errors := &LoginUserErrors{}

	// TODO: validate through validator pkg
	return errors
}

func (errors *LoginUserErrors) SetOtherError(message string) {
	errors.Other = message
}
