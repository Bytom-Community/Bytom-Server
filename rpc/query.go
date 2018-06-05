package rpc

import (
	"context"
	"fmt"

	"github.com/bytom/blockchain/query"
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

func (s *ApiService) ListBalances(ctx context.Context, req *rpcpb.ListBalancesRequest) (*rpcpb.ListBalancesResponse, error) {
	accountID := req.AccountID
	balances, err := s.wallet.GetAccountBalances(accountID)
	if err != nil {
		return nil, fmt.Errorf("list-balances: %v", err.Error())
	}

	var results []string
	for _, balance := range balances {
		results = append(results, string(util.JsonEncode(balance)))
	}

	return &rpcpb.ListBalancesResponse{Balances: results}, nil
}

func (s *ApiService) ListTransactions(ctx context.Context, req *rpcpb.ListTransactionsRequest) (*rpcpb.ListTransactionsResponse, error) {
	txID := req.TxID
	accountID := req.AccountID
	detail := req.Detail
	transactions := []*query.AnnotatedTx{}
	var err error
	var results []string

	if accountID != "" {
		transactions, err = s.wallet.GetTransactionsByAccountID(accountID)
	} else {
		transactions, err = s.wallet.GetTransactionsByTxID(txID)
	}

	if err != nil {
		return nil, fmt.Errorf("list-transactions: %v", err.Error())
	}

	if detail == false {
		txSummaries := s.wallet.GetTransactionsSummary(transactions)
		for _, txSum := range txSummaries {
			results = append(results, string(util.JsonEncode(txSum)))
		}
		return &rpcpb.ListTransactionsResponse{Transactions: results}, nil
	}

	for _, trans := range transactions {
		results = append(results, string(util.JsonEncode(trans)))
	}
	return &rpcpb.ListTransactionsResponse{Transactions: results}, nil
}
