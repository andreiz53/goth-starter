package server

import (
	"fmt"
	"goth/cookies"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	listenAddr  string
	cookieStore cookies.CookieStore
}

func NewServer(addr string, cs cookies.CookieStore) Server {
	return Server{
		listenAddr:  addr,
		cookieStore: cs,
	}
}

func (s Server) Init() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(s.WithUser)

	fs := http.FileServer(http.Dir("web/assets"))
	router.Handle("/assets/*", http.StripPrefix("/assets", fs))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) })

	router.Get("/", s.renderIndex)
	router.Get("/login", s.renderLogin)
	router.Get("/register", s.renderRegister)

	router.Post("/login", s.handleLogin)
	router.Post("/register", s.handleRegister)
	router.Post("/logout", s.handleLogout)

	fmt.Printf("Server starting on port %s\n", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
