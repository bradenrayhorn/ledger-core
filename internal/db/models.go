// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"github.com/google/uuid"
)

type UserMarketProvider struct {
	UserUuid uuid.UUID
	Provider string
}
