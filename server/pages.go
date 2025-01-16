package server

import (
	"goth/web/pages"
	"net/http"
)

func (s Server) renderIndex(w http.ResponseWriter, r *http.Request) {
	RenderComponent(w, r, pages.Index())
}

func (s Server) renderLogin(w http.ResponseWriter, r *http.Request) {}

func (s Server) renderRegister(w http.ResponseWriter, r *http.Request) {}
