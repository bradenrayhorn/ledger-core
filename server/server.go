package server

import (
	"log"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/bradenrayhorn/ledger-core/http"
)

type server struct {
	Config     *core.Config
	httpServer *http.Server
}

func CreateServer() *server {
	config := loadConfig()

	return &server{
		Config:     config,
		httpServer: http.CreateServer(config),
	}
}

func (s server) GetHttpServer() *http.Server {
	return s.httpServer
}

func (s server) Run() error {
	log.Println("initializing ledger-core...")

	s.httpServer.Start()
	return nil
}
