package rpc

import (
	"context"

	"github.com/Bytom-Community/bytom-wallet/rpc/pb"
)

type ApiService struct {
	RpcServer
}

func (s *ApiService)CreateKey(ctx context.Context, req *rpcpb.CreateKeyRequest) (*rpcpb.CreateKeyResponse, error) {

}

func (s *ApiService)ListKey(ctx context.Context, req *rpcpb.NonParamsRequest) (*rpcpb.ListKeyResponse, error) {

}