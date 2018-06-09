package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/bytom/blockchain/txbuilder"
	"github.com/bytom/consensus"
	"github.com/bytom/consensus/segwit"
	"github.com/bytom/errors"
	"github.com/bytom/math/checked"
	"github.com/bytom/rpc/pb"
)

var (
	defaultTxTTL    = 5 * time.Minute
	defaultBaseRate = float64(100000)
)

func (s *ApiService) SubmitTransaction(ctx context.Context, req *rpcpb.SubmitTransactionRequest) (*rpcpb.SubmitTransactionResponse, error) {
	// rawTx := new(types.Tx)
	// if err := rawTx.UnmarshalText([]byte(req.RawTransaction)); err != nil {
	// 	return nil, fmt.Errorf("submit-transaction: %v", err.Error())
	// }
	// if err := txbuilder.FinalizeTx(ctx, s.chain, rawTx); err != nil {
	// 	return nil, fmt.Errorf("submit-transaction: %v", err.Error())
	// }
	// return &rpcpb.SubmitTransactionResponse{TxID: rawTx.ID.String()}, nil
	return nil, nil
}

func (s *ApiService) EstimateTransactionGas(ctx context.Context, req *rpcpb.EstimateTransactionGasRequest) (*rpcpb.EstimateTransactionGasResponse, error) {
	txTemplate := new(txbuilder.Template)
	if err := json.Unmarshal([]byte(req.TxTemplate), txTemplate); err != nil {
		return nil, fmt.Errorf("estimate-transaction-gas: %v", err.Error())
	}

	txGasResp, err := EstimateTxGas(*txTemplate)
	if err != nil {
		return nil, fmt.Errorf("estimate-transaction-gas: %v", err.Error())
	}

	return &rpcpb.EstimateTransactionGasResponse{
		TotalNeu:   txGasResp.TotalNeu,
		StorageNeu: txGasResp.StorageNeu,
		VmNeu:      txGasResp.VMNeu,
	}, nil
}

// EstimateTxGasResp estimate transaction consumed gas
type EstimateTxGasResp struct {
	TotalNeu   int64 `json:"total_neu"`
	StorageNeu int64 `json:"storage_neu"`
	VMNeu      int64 `json:"vm_neu"`
}

// EstimateTxGas estimate consumed neu for transaction
func EstimateTxGas(template txbuilder.Template) (*EstimateTxGasResp, error) {
	// base tx size and not include sign
	data, err := template.Transaction.TxData.MarshalText()
	if err != nil {
		return nil, err
	}
	baseTxSize := int64(len(data))

	// extra tx size for sign witness parts
	baseWitnessSize := int64(300)
	lenSignInst := int64(len(template.SigningInstructions))
	signSize := baseWitnessSize * lenSignInst

	// total gas for tx storage
	totalTxSizeGas, ok := checked.MulInt64(baseTxSize+signSize, consensus.StorageGasRate)
	if !ok {
		return nil, errors.New("calculate txsize gas got a math error")
	}

	// consume gas for run VM
	totalP2WPKHGas := int64(0)
	totalP2WSHGas := int64(0)
	baseP2WPKHGas := int64(1419)
	baseP2WSHGas := int64(2499)

	for _, inpID := range template.Transaction.Tx.InputIDs {
		sp, err := template.Transaction.Spend(inpID)
		if err != nil {
			continue
		}

		resOut, err := template.Transaction.Output(*sp.SpentOutputId)
		if err != nil {
			continue
		}

		if segwit.IsP2WPKHScript(resOut.ControlProgram.Code) {
			totalP2WPKHGas += baseP2WPKHGas
		} else if segwit.IsP2WSHScript(resOut.ControlProgram.Code) {
			totalP2WSHGas += baseP2WSHGas
		}
	}

	// total estimate gas
	totalGas := totalTxSizeGas + totalP2WPKHGas + totalP2WSHGas

	// rounding totalNeu with base rate 100000
	totalNeu := float64(totalGas*consensus.VMGasRate) / defaultBaseRate
	roundingNeu := math.Ceil(totalNeu)
	estimateNeu := int64(roundingNeu) * int64(defaultBaseRate)

	return &EstimateTxGasResp{
		TotalNeu:   estimateNeu,
		StorageNeu: totalTxSizeGas * consensus.VMGasRate,
		VMNeu:      (totalP2WPKHGas + totalP2WSHGas) * consensus.VMGasRate,
	}, nil
}
