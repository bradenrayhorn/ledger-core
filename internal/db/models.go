// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"

	"github.com/google/uuid"
)

type UserMarketProvider struct {
	UserUuid  uuid.UUID
	Provider  string
	CreatedAt time.Time
}
