package server

import (
	"log"

	core "github.com/bradenrayhorn/ledger-core"
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

	return nil
}
