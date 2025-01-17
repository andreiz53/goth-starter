package server

import (
	"fmt"
	"goth/types"
	"goth/web/pages"
	"net/http"
)

func (s Server) renderIndex(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(types.UserContextKey).(types.AuthenticatedUser)
	fmt.Println("got user", user, "ok", ok)
	RenderComponent(w, r, pages.Index())
}

func (s Server) renderLogin(w http.ResponseWriter, r *http.Request) {
	RenderComponent(w, r, pages.Login())
}

func (s Server) renderRegister(w http.ResponseWriter, r *http.Request) {
	RenderComponent(w, r, pages.Register())
}
