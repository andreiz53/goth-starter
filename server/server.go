package server

import (
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

	http.ListenAndServe(s.ListenAddr, router)
}
