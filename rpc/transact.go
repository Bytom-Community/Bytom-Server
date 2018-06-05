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

func (s *ApiService) SubmitTranstraction(ctx context.Context, req *rpcpb.SubmitTranstractionRequest) (*rpcpb.SubmitTranstractionResponse, error) {
	rawTx := new(types.Tx)
	if err := rawTx.UnmarshalText([]byte(req.RawTranstraction)); err != nil {
		return nil, fmt.Errorf("submit-transtraction: %v", err.Error())
	}
	if err := txbuilder.FinalizeTx(ctx, s.chain, rawTx); err != nil {
		return nil, fmt.Errorf("submit-transtraction: %v", err.Error())
	}
	return &rpcpb.SubmitTranstractionResponse{TxID: rawTx.ID}, nil
}
