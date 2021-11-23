package server

import (
	core "github.com/bradenrayhorn/ledger-core"
	"github.com/bradenrayhorn/ledger-core/grpc"
	"github.com/bradenrayhorn/ledger-core/http"
	"github.com/bradenrayhorn/ledger-core/internal/db"
	"github.com/bradenrayhorn/ledger-core/postgres"
	"github.com/bradenrayhorn/ledger-core/zap"
	"github.com/jackc/pgx/v4/pgxpool"
)

type server struct {
	Config     *core.Config
	httpServer *http.Server
	pgxPool    *pgxpool.Pool
	logger     core.Logger
}

func CreateServer(config *core.Config) *server {
	return &server{
		Config: config,
		logger: zap.CreateLogger(config),
	}
}

func (s server) GetHttpServer() *http.Server {
	return s.httpServer
}

func (s *server) Setup() error {
	s.logger.Info("initializing ledger-core...")

	pool, err := postgres.CreatePool(s.Config)
	if err != nil {
		return err
	}
	s.pgxPool = pool

	s.httpServer = &http.Server{
		Config:                 s.Config,
		UserMarketProviderRepo: postgres.NewUserMarketProviderRepository(db.New(s.pgxPool)),
		SessionService:         grpc.NewSessionService(s.Config.GrpcConn),
		Logger:                 s.logger,
	}
	s.httpServer.Initialize()

	return nil
}

func (s *server) Run() {
	s.logger.Info("starting ledger-core...")

	s.httpServer.Start()
}
