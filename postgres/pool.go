package postgres

import (
	"context"
	"fmt"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreatePool(config *core.Config) (*pgxpool.Pool, error) {
	pgxConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s%s",
		config.PgUsername,
		config.PgPassword,
		config.PgHost,
		config.PgPort,
		config.PgDatabase,
		config.PgParameters,
	))
	if err != nil {
		return nil, err
	}

	return pgxpool.ConnectConfig(context.Background(), pgxConfig)
}
