package storage

import (
	"os"
	"testing"

	"github.com/DavudSafarli/go-web-starter-template/domains/appname/contracts"
	"github.com/stretchr/testify/require"
)

func TestPostgres(t *testing.T) {
	connstr := os.Getenv("TEST_POSTGRES_URL")
	if connstr == "" {
		t.Skip(`Skipping TestPostgres because "TEST_POSTGRES_URL" is not set`)
	}

	pg, err := NewPostgres(connstr)
	require.Nil(t, err)

	contracts.StorageContract{
		Subject: pg,
	}.Test(t)
}
