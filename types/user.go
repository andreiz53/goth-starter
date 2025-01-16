package types

type User struct {
	Email    string
	Password string
}

type AuthenticatedUser struct {
	Email       string
	AccessToken string
}
