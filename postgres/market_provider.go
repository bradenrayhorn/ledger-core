package postgres

import (
	"context"
	"errors"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/bradenrayhorn/ledger-core/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type UserMarketProviderRepository struct {
	db *db.Queries
}

func NewUserMarketProviderRepository(db *db.Queries) *UserMarketProviderRepository {
	return &UserMarketProviderRepository{
		db: db,
	}
}

func (r *UserMarketProviderRepository) GetUserMarketProvider(ctx context.Context, userUUID uuid.UUID) (*core.UserMarketProvider, error) {
	userMarketProvider, err := r.db.GetUserMarketProvider(ctx, userUUID)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &core.UserMarketProvider{
		UserUUID: userMarketProvider.UserUuid,
		Provider: userMarketProvider.Provider,
	}, nil
}

func (r *UserMarketProviderRepository) SetUserMarketProvider(ctx context.Context, userUUID uuid.UUID, provider string) error {
	return r.db.SetUserMarketProvider(ctx, db.SetUserMarketProviderParams{
		UserUuid: userUUID,
		Provider: provider,
	})
}

func (r *UserMarketProviderRepository) DeleteUserMarketProvider(ctx context.Context, userUUID uuid.UUID) error {
	return nil
}
