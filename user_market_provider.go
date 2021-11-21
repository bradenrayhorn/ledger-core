package core

import (
	"context"

	"github.com/google/uuid"
)

type UserMarketProvider struct {
	UserUUID uuid.UUID
	Provider string
}

type UserMarketProviderRepository interface {
	GetUserMarketProvider(ctx context.Context, userUUID uuid.UUID) (*UserMarketProvider, error)
	SetUserMarketProvider(ctx context.Context, userUUID uuid.UUID, provider string) error
	DeleteUserMarketProvider(ctx context.Context, userUUID uuid.UUID) error
}
