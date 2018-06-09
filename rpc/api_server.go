package rpc

import (
	"context"

	"github.com/bytom/chaincache"
	"github.com/bytom/rpc/pb"
)

type ApiService struct {
	rpcServer *Rpc

	chainCache *chaincache.ChainCache
}

func (s *ApiService) GetState(ctx context.Context, req *rpcpb.NonParamsRequest) (*rpcpb.GetStateResponse, error) {
	return &rpcpb.GetStateResponse{Status: "OK"}, nil
}
