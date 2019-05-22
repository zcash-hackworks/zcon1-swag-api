package swag

import (
	"context"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/zcash-hackworks/zcon1-swag-api/api"
)

type Server struct {
}

func NewServer(dbPath string, zc *rpcclient.Client) (api.SwagAPIServer, error) {
	return &Server{}, nil
}

func (s *Server) Redeem(ctx context.Context, req *api.Request) (*api.Response, error) {
	return &api.Response{
		Success: true,
		Msg:     "hello",
	}, nil
}

func (s *Server) GracefulStop() {}
