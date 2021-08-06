package server

import (
	"log"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/bradenrayhorn/ledger-core/http"
)

type server struct {
	config *core.Config
}

func CreateServer() *server {
	config := loadConfig()

	return &server{
		config: config,
	}
}

func (s server) Run() error {
	log.Println("initializing ledger-core...")

	httpServer := http.CreateServer(s.config)
	httpServer.Start()

	return nil
}
