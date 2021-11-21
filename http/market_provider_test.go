package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradenrayhorn/ledger-core/internal/testutils"
	"github.com/bradenrayhorn/ledger-protos/session"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCanGetMarketProvider(t *testing.T) {
	sv, grpcSv := testutils.SetupHTTPServer()
	defer grpcSv.Close()

	grpcSv.SessionService.On("Authenticate", mock.Anything, mock.MatchedBy(func(r *session.SessionAuthenticateRequest) bool { return r.GetSessionID() == "my-session" })).Return(&session.SessionAuthenticateResponse{
		Session: &session.Session{
			SessionID: "my-session",
			UserID:    uuid.New().String(),
		},
	}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/market-provider", nil)
	req.Header.Add("cookie", "session_id=my-session")
	res := httptest.NewRecorder()
	sv.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.JSONEq(t, `{"data": null}`, res.Body.String())
}
