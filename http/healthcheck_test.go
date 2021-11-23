package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradenrayhorn/ledger-core/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	sv := testutils.SetupHTTPServer()
	defer sv.Grpc.Close()

	req, _ := http.NewRequest(http.MethodGet, "/health-check", nil)
	res := httptest.NewRecorder()
	sv.Http.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
}
