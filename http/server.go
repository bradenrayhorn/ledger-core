package http

import (
	"net/http"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type server struct {
	config *core.Config
}

func CreateServer(config *core.Config) *server {
	return &server{
		config: config,
	}
}

func (s *server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health-check", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":"+s.config.HttpPort, r)
}
