package rpc

import (
	"context"
	"github.com/bytom/rpc/pb"
)

func (s *ApiService) ListAssets(ctx context.Context, req *rpcpb.ListAssetsRequest) (*rpcpb.ListAssetsResponse, error) {
	var assets []*rpcpb.Asset
	for assetID, amount := range s.chainCache.ListAssets(req.Address) {
		asset := &rpcpb.Asset{
			AssetID: assetID,
			Amount:  amount,
		}
		assets = append(assets, asset)
	}
	return &rpcpb.ListAssetsResponse{Assets: assets}, nil
}

func (s *ApiService) ListTransactions(ctx context.Context, req *rpcpb.ListTransactionsRequest) (*rpcpb.ListTransactionsResponse, error) {
	var transactions []*rpcpb.TX
	for _, tx := range s.chainCache.ListTransactions(req.Address, req.AssetID) {
		var inputs []*rpcpb.Input
		var outputs []*rpcpb.Output
		var op string

		for _, v := range tx.Inputs {
			input := &rpcpb.Input{
				Type:          v.Type,
				AssetID:       v.AssetID.String(),
				Amount:        v.Amount,
				Address:       v.Address,
				SpentOutputID: v.SpentOutputID.String(),
			}
			if input.Address == req.Address {
				op = "send"
			}
			inputs = append(inputs, input)
		}

		for _, v := range tx.Outputs {
			output := &rpcpb.Output{
				Type:     v.Type,
				AssetID:  v.AssetID.String(),
				Amount:   v.Amount,
				Address:  v.Address,
				OutputID: v.OutputID.String(),
				Position: int32(v.Position),
			}
			if output.Address == req.Address {
				op = "receive"
			}
			outputs = append(outputs, output)
		}
		TX := &rpcpb.TX{
			ID:                     tx.ID.String(),
			Timestamp:              tx.Timestamp,
			BlockID:                tx.BlockID.String(),
			BlockHeight:            tx.BlockHeight,
			Position:               tx.Position,
			BlockTransactionsCount: tx.BlockTransactionsCount,
			Confirmation:           s.chainCache.BestBlockHeight() - tx.BlockHeight,
			StatusFail:             tx.StatusFail,
			Inputs:                 inputs,
			Outputs:                outputs,
			Op:                     op,
		}
		transactions = append(transactions, TX)
	}
	return &rpcpb.ListTransactionsResponse{Transactions: transactions}, nil
}
