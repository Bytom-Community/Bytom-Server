package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/bytom/blockchain/txbuilder"
	"github.com/bytom/protocol/bc/types"
	"github.com/bytom/rpc/pb"
)

var (
	defaultTxTTL    = 5 * time.Minute
	defaultBaseRate = float64(100000)
)

func (s *ApiService) SubmitTransaction(ctx context.Context, req *rpcpb.SubmitTransactionRequest) (*rpcpb.SubmitTransactionResponse, error) {
	rawTx := new(types.Tx)
	if err := rawTx.UnmarshalText([]byte(req.RawTranstraction)); err != nil {
		return nil, fmt.Errorf("submit-transaction: %v", err.Error())
	}
	if err := txbuilder.FinalizeTx(ctx, s.chain, rawTx); err != nil {
		return nil, fmt.Errorf("submit-transaction: %v", err.Error())
	}
	return &rpcpb.SubmitTransactionResponse{TxID: rawTx.ID}, nil
}
