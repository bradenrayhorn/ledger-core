package testutils

import (
	"log"
	"os"
	"strings"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/bradenrayhorn/ledger-core/server"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func SetupHTTPServer() (*chi.Mux, *MockGrpcServer) {
	mockGRPC := SetupGrpcConn()

	loadConfig()

	config := &core.Config{
		HttpPort:     "80",
		PgHost:       viper.GetString("pg_host"),
		PgPort:       viper.GetString("pg_port"),
		PgUsername:   viper.GetString("pg_username"),
		PgPassword:   viper.GetString("pg_password"),
		PgDatabase:   viper.GetString("pg_database"),
		PgParameters: viper.GetString("pg_parameters"),

		GrpcConn: mockGRPC.Conn,
	}

	sv := server.CreateServer(config)
	sv.Setup()

	return sv.GetHttpServer().GetRouter(), mockGRPC
}

func loadConfig() {
	envPath := os.Getenv("TEST_ENV_PATH")
	if len(envPath) == 0 {
		envPath = "../.env.test"
	}

	if _, err := os.Stat(envPath); err == nil {
		err := godotenv.Load(envPath)
		if err != nil {
			log.Printf("failed to load .env: %s", err)
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
