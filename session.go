package core

import "context"

type SessionService interface {
	GetUserFromSession(ctx context.Context, sessionID string, ip string, userAgent string) (string, error)
}
