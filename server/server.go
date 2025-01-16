package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	ListenAddr string
}

func NewServer(addr string) Server {
	return Server{
		ListenAddr: addr,
	}
}

func (s Server) Init() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("web/assets"))
	router.Handle("/assets/*", http.StripPrefix("/assets", fs))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) })

	router.Get("/", s.renderIndex)
	router.Get("/login", s.renderLogin)
	router.Get("/register", s.renderRegister)

	router.Post("/login", s.handleLogin)
	router.Post("/register", s.handleRegister)
	router.Post("/logout", s.handleLogout)

	fmt.Printf("Server starting on port %s\n", s.ListenAddr)
	http.ListenAndServe(s.ListenAddr, router)
}
