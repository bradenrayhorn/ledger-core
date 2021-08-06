package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/bradenrayhorn/ledger-core/server"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

var router http.Handler

func testMain(m *testing.M) int {
	server := server.CreateServer()
	router = server.GetHttpServer().GetRouter()
	return m.Run()
}
