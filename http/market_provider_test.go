package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradenrayhorn/ledger-core/internal/db"
	"github.com/bradenrayhorn/ledger-core/internal/testutils"
	"github.com/bradenrayhorn/ledger-core/postgres"
	"github.com/stretchr/testify/assert"
)

func TestCanGetBlankMarketProvider(t *testing.T) {
	sv := testutils.SetupHTTPServer()
	defer sv.Grpc.Close()

	_ = testutils.FakeAuthedUser(sv)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/market-provider", nil)
	req.Header.Add("cookie", "session_id=my-session")
	res := httptest.NewRecorder()
	sv.Http.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.JSONEq(t, `{"data": null}`, res.Body.String())
}

func TestCanGetMarketProvider(t *testing.T) {
	sv := testutils.SetupHTTPServer()
	defer sv.Grpc.Close()

	userID := testutils.FakeAuthedUser(sv)
	userRepo := postgres.NewUserMarketProviderRepository(db.New(sv.Pg))
	userRepo.SetUserMarketProvider(context.Background(), userID, "tda")

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/market-provider", nil)
	req.Header.Add("cookie", "session_id=my-session")
	res := httptest.NewRecorder()
	sv.Http.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.JSONEq(t, `{"data": "tda"}`, res.Body.String())
}
