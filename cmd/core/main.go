package main

import (
	"log"

	"github.com/bradenrayhorn/ledger-core/server"
)

func main() {
	server := server.CreateServer()
	err := server.Run()

	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
