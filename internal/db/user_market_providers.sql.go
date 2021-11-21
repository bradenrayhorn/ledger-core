// Code generated by sqlc. DO NOT EDIT.
// source: user_market_providers.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUserMarketProvider = `-- name: CreateUserMarketProvider :exec
INSERT INTO user_market_providers (
  user_uuid, provider
) VALUES ($1, $2)
`

type CreateUserMarketProviderParams struct {
	UserUuid uuid.UUID
	Provider string
}

func (q *Queries) CreateUserMarketProvider(ctx context.Context, arg CreateUserMarketProviderParams) error {
	_, err := q.db.Exec(ctx, createUserMarketProvider, arg.UserUuid, arg.Provider)
	return err
}

const getUserMarketProvider = `-- name: GetUserMarketProvider :one
SELECT user_uuid, provider, created_at FROM user_market_providers WHERE user_uuid = $1
`

func (q *Queries) GetUserMarketProvider(ctx context.Context, userUuid uuid.UUID) (UserMarketProvider, error) {
	row := q.db.QueryRow(ctx, getUserMarketProvider, userUuid)
	var i UserMarketProvider
	err := row.Scan(&i.UserUuid, &i.Provider, &i.CreatedAt)
	return i, err
}