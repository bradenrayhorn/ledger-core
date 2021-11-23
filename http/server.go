package http

import (
	"fmt"
	"net/http"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Server struct {
	Config                 *core.Config
	Logger                 core.Logger
	UserMarketProviderRepo core.UserMarketProviderRepository
	SessionService         core.SessionService

	router *chi.Mux
}

func (s *Server) Initialize() {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON), middleware.Logger)
	r.Get("/health-check", HealthCheck)

	r.Route("/api", func(r chi.Router) {
		r.Use(render.SetContentType(render.ContentTypeJSON), s.authentication())

		r.Route("/v1", func(r chi.Router) {
			r.Route("/market-provider", func(r chi.Router) {
				r.Get("/", s.GetMarketProvider)
				r.Post("/", s.UpdateMarketProvider)
			})
		})
	})

	s.router = r
}

func (s *Server) GetRouter() *chi.Mux {
	return s.router
}

func (s *Server) Start() {
	s.Logger.Info(fmt.Sprintf("starting http server on port %s", s.Config.HttpPort))

	http.ListenAndServe(":"+s.Config.HttpPort, s.router)
}
