package rpc

import (
	"context"
	// "fmt"
	// "github.com/bytom/blockchain/query"
	"github.com/bytom/rpc/pb"
	"github.com/bytom/util"
	// "github.com/bytom/util"
)

func (s *ApiService) ListAssets(ctx context.Context, req *rpcpb.ListAssetsRequest) (*rpcpb.ListAssetsResponse, error) {
	address := req.Address
	assets := s.chainCache.ListAssets(address)
	results := string(util.JsonEncode(assets))
	return &rpcpb.ListAssetsResponse{Assets: results}, nil
}

func (s *ApiService) ListBalances(ctx context.Context, req *rpcpb.ListBalancesRequest) (*rpcpb.ListBalancesResponse, error) {
	// accountID := req.AccountID
	// balances, err := s.wallet.GetAccountBalances(accountID)
	// if err != nil {
	// 	return nil, fmt.Errorf("list-balances: %v", err.Error())
	// }

	// var results []string
	// for _, balance := range balances {
	// 	results = append(results, string(util.JsonEncode(balance)))
	// }

	// return &rpcpb.ListBalancesResponse{Balances: results}, nil
	return nil, nil
}

// ListTransactions list transactions form chain
func (s *ApiService) ListTransactions(ctx context.Context, req *rpcpb.ListTransactionsRequest) (*rpcpb.ListTransactionsResponse, error) {
	address := req.Address
	assetID := req.AssetID

	txs := s.chainCache.ListTransactions(address, assetID)

	var transactions []string
	for _, tx := range txs {
		transactions = append(transactions, string(util.JsonEncode(tx)))
	}
	return &rpcpb.ListTransactionsResponse{Transactions: transactions}, nil
}

func (s *ApiService) ListTransaction(ctx context.Context, req *rpcpb.ListTransactionRequest) (*rpcpb.ListTransactionResponse, error) {
	txID := req.TxID
	tx := s.chainCache.ListTransaction(txID)

	return &rpcpb.ListTransactionResponse{Tx: string(util.JsonEncode(tx))}, nil
}
