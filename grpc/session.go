package grpc

import (
	"context"

	"github.com/bradenrayhorn/ledger-protos/session"
	"google.golang.org/grpc"
)

type SessionService struct {
	client session.SessionAuthenticatorClient
}

func NewSessionService(grpcConn *grpc.ClientConn) *SessionService {
	return &SessionService{
		client: session.NewSessionAuthenticatorClient(grpcConn),
	}
}

func (s *SessionService) GetUserFromSession(ctx context.Context, sessionID string, ip string, userAgent string) (string, error) {
	res, err := s.client.Authenticate(ctx, &session.SessionAuthenticateRequest{
		SessionID: sessionID,
		UserAgent: userAgent,
		IP:        ip,
	})

	if err != nil {
		return "", err
	}

	return res.Session.GetUserID(), nil
}
