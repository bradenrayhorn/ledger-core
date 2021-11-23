package testutils

import (
	"log"
	"os"
	"strings"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/bradenrayhorn/ledger-core/postgres"
	"github.com/bradenrayhorn/ledger-core/server"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type TestServer struct {
	Http *chi.Mux
	Grpc *MockGrpcServer
	Pg   *pgxpool.Pool
}

func SetupHTTPServer() TestServer {
	mockGRPC := SetupGrpcConn()

	loadConfig()

	config := &core.Config{
		HttpPort: "80",

		LogLevel:  core.LogLevelDebug,
		LogFormat: core.LogFormatConsole,

		PgHost:       viper.GetString("pg_host"),
		PgPort:       viper.GetString("pg_port"),
		PgUsername:   viper.GetString("pg_username"),
		PgPassword:   viper.GetString("pg_password"),
		PgDatabase:   viper.GetString("pg_database"),
		PgParameters: viper.GetString("pg_parameters"),

		GrpcConn: mockGRPC.Conn,
	}

	sv := server.CreateServer(config)
	err := sv.Setup()
	if err != nil {
		log.Fatalf("server setup failed, %v", err)
	}

	pg, _ := postgres.CreatePool(config)

	return TestServer{
		Http: sv.GetHttpServer().GetRouter(),
		Grpc: mockGRPC,
		Pg:   pg,
	}
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
