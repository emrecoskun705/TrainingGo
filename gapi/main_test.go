package gapi

import (
	"context"
	"fmt"
	db "github.com/emrecoskun705/TraniningGo/db/sqlc"
	"github.com/emrecoskun705/TraniningGo/token"
	"github.com/emrecoskun705/TraniningGo/util"
	"github.com/emrecoskun705/TraniningGo/worker"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
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

func newContextWithErrorToken(t *testing.T, tokenMaker token.Maker, username string, duration time.Duration) context.Context {
	accessToken, _, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)
	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	md := metadata.MD{
		authorizationHeader: []string{
			bearerToken,
		},
	}
	return metadata.NewIncomingContext(context.Background(), md)
}