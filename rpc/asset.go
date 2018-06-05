package rpc

import (
	"context"
	"fmt"

	"github.com/bytom/rpc/pb"
	"github.com/bytom/util"
)

func (s *ApiService) ListAssets(ctx context.Context, req *rpcpb.ListAssetsRequest) (*rpcpb.ListAssetsResponse, error) {
	assetID := req.AssetID
	assets, err := s.wallet.AssetReg.ListAssets(assetID)
	if err != nil {
		return nil, fmt.Errorf("list-assets: %v", err.Error())
	}

	var results []string
	for _, asset := range assets {
		results = append(results, string(util.JsonEncode(asset)))
	}

	return &rpcpb.ListAssetsResponse{Assets: results}, nil
}
