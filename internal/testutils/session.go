package testutils

import (
	"github.com/bradenrayhorn/ledger-protos/session"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func FakeAuthedUser(sv TestServer) uuid.UUID {
	userUUID := uuid.New()

	sv.Grpc.SessionService.On("Authenticate", mock.Anything, mock.MatchedBy(func(r *session.SessionAuthenticateRequest) bool { return r.GetSessionID() == "my-session" })).Return(&session.SessionAuthenticateResponse{
		Session: &session.Session{
			SessionID: "my-session",
			UserID:    userUUID.String(),
		},
	}, nil)

	return userUUID
}
