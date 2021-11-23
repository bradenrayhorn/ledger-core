package http_test

import (
	"context"
	"os"
	"testing"

	"github.com/bradenrayhorn/ledger-core/internal/testutils"
)

func TestMain(m *testing.M) {
	sv := testutils.SetupHTTPServer()
	sv.Pg.Exec(context.Background(), "truncate table user_market_providers;")

	os.Exit(m.Run())
}
