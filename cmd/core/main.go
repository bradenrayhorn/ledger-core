package main

import (
	"log"

	"github.com/bradenrayhorn/ledger-core/config"
	"github.com/bradenrayhorn/ledger-core/server"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("invalid config")
		return
	}

	server := server.CreateServer(config)
	err = server.Setup()
	if err != nil {
		log.Fatalf("failed to setup server: %s", err)
	}
	server.Run()

	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
