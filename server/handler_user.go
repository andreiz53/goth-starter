package server

import (
	"goth/types"
	"goth/utils"
	"goth/web/components/forms"
	"net/http"
)

func (s Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	values, errors := parseAndValidateForm[types.RegisterUserValues](r)
	if *errors != (types.RegisterUserErrors{}) {
		RenderComponent(w, r, forms.RegisterForm(values, *errors))
		return
	}

	// encrypt password
	encPw := utils.EncryptPassword(values.Password)
	// insert into db
	_ = encPw

	// set user in cookie store
	user := types.AuthenticatedUser{
		Email:    values.Email,
		LoggedIn: true,
	}
	s.loginUser(w, r, user)
	Redirect(w, r, "/")
}

func (s Server) handleLogin(w http.ResponseWriter, r *http.Request) {}

func (s Server) handleLogout(w http.ResponseWriter, r *http.Request) {
	s.logoutUser(w, r)
	Redirect(w, r, "/login")
}
