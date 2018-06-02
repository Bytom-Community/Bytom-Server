package rpc

import (
	"context"

	"github.com/bytom/protocol"
	"github.com/bytom/wallet"
)

type ApiService struct {
	rpcServer *Rpc

	wallet *wallet.Wallet
	chain  *protocol.Chain
}

func (s *ApiService) GetState(ctx context.Context, req *rpcpb.NonParamsRequest) (*rpcpb.GetStateResponse, error) {
	return &rpcpb.GetStateResponse{Status: "OK"}, nil
}
