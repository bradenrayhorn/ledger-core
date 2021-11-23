package http_test

import (
	"bytes"
	"context"
	"encoding/json"
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

	res := getMarketProvider(sv)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.JSONEq(t, `{"data": null}`, res.Body.String())
}

func TestCanGetMarketProvider(t *testing.T) {
	sv := testutils.SetupHTTPServer()
	defer sv.Grpc.Close()

	userID := testutils.FakeAuthedUser(sv)
	userRepo := postgres.NewUserMarketProviderRepository(db.New(sv.Pg))
	userRepo.SetUserMarketProvider(context.Background(), userID, "tda")

	res := getMarketProvider(sv)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.JSONEq(t, `{"data": "tda"}`, res.Body.String())
}

func TestCanSetMarketProvider(t *testing.T) {
	sv := testutils.SetupHTTPServer()
	defer sv.Grpc.Close()

	_ = testutils.FakeAuthedUser(sv)

	res := setMarketProvider(sv, "tda")
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)

	res = getMarketProvider(sv)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.JSONEq(t, `{"data": "tda"}`, res.Body.String())
}

func TestCanUpdateMarketProvider(t *testing.T) {
	sv := testutils.SetupHTTPServer()
	defer sv.Grpc.Close()

	userID := testutils.FakeAuthedUser(sv)

	userRepo := postgres.NewUserMarketProviderRepository(db.New(sv.Pg))
	userRepo.SetUserMarketProvider(context.Background(), userID, "tda-two")

	res := setMarketProvider(sv, "tda")
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)

	res = getMarketProvider(sv)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.JSONEq(t, `{"data": "tda"}`, res.Body.String())
}

func TestCanClearMarketProvider(t *testing.T) {
	sv := testutils.SetupHTTPServer()
	defer sv.Grpc.Close()

	userID := testutils.FakeAuthedUser(sv)

	userRepo := postgres.NewUserMarketProviderRepository(db.New(sv.Pg))
	userRepo.SetUserMarketProvider(context.Background(), userID, "tda")

	res := setMarketProvider(sv, "")
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)

	res = getMarketProvider(sv)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	assert.JSONEq(t, `{"data": null}`, res.Body.String())
}

func getMarketProvider(sv testutils.TestServer) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/market-provider", nil)
	req.Header.Add("cookie", "session_id=my-session")
	res := httptest.NewRecorder()
	sv.Http.ServeHTTP(res, req)
	return res
}

func setMarketProvider(sv testutils.TestServer, provider string) *httptest.ResponseRecorder {
	json, _ := json.Marshal(map[string]string{"provider": provider})
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/market-provider", bytes.NewReader(json))
	req.Header.Add("cookie", "session_id=my-session")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()
	sv.Http.ServeHTTP(res, req)
	return res
}
