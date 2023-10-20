package gapi

import (
	db "github.com/emrecoskun705/TraniningGo/db/sqlc"
	"github.com/emrecoskun705/TraniningGo/util"
	"github.com/emrecoskun705/TraniningGo/worker"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
