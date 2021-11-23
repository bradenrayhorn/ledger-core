// Code generated by sqlc. DO NOT EDIT.
// source: user_market_providers.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getUserMarketProvider = `-- name: GetUserMarketProvider :one
SELECT user_uuid, provider FROM user_market_providers WHERE user_uuid = $1
`

func (q *Queries) GetUserMarketProvider(ctx context.Context, userUuid uuid.UUID) (UserMarketProvider, error) {
	row := q.db.QueryRow(ctx, getUserMarketProvider, userUuid)
	var i UserMarketProvider
	err := row.Scan(&i.UserUuid, &i.Provider)
	return i, err
}

const setUserMarketProvider = `-- name: SetUserMarketProvider :exec
INSERT INTO user_market_providers (
  user_uuid, provider
) VALUES ($1, $2) ON CONFLICT(user_uuid) DO UPDATE SET provider = $2
`

type SetUserMarketProviderParams struct {
	UserUuid uuid.UUID
	Provider string
}

func (q *Queries) SetUserMarketProvider(ctx context.Context, arg SetUserMarketProviderParams) error {
	_, err := q.db.Exec(ctx, setUserMarketProvider, arg.UserUuid, arg.Provider)
	return err
}
