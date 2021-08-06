package http

import (
	"net/http"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	config *core.Config
	router *chi.Mux
}

func CreateServer(config *core.Config) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health-check", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
	})

	return &Server{
		config: config,
		router: r,
	}
}

func (s *Server) GetRouter() *chi.Mux {
	return s.router
}

func (s *Server) Start() {
	http.ListenAndServe(":"+s.config.HttpPort, s.router)
}
