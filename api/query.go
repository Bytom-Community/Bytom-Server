package api

import (
	"context"
	"fmt"
	"sort"

	log "github.com/sirupsen/logrus"

	"github.com/bytom/account"
	"github.com/bytom/blockchain/query"
	"github.com/bytom/consensus"
	"github.com/bytom/db"
	chainjson "github.com/bytom/encoding/json"
	"github.com/bytom/errors"
	"github.com/bytom/protocol/bc"
	"github.com/bytom/protocol/bc/types"
)

// POST /list-accounts
func (a *API) listAccounts(ctx context.Context, filter struct {
	ID string `json:"id"`
}) Response {
	accounts, err := a.wallet.AccountMgr.ListAccounts(filter.ID)
	if err != nil {
		log.Errorf("listAccounts: %v", err)
		return NewErrorResponse(err)
	}

	annotatedAccounts := []query.AnnotatedAccount{}
	for _, acc := range accounts {
		annotatedAccounts = append(annotatedAccounts, *account.Annotated(acc))
	}

	return NewSuccessResponse(annotatedAccounts)
}

// POST /get-asset
func (a *API) getAsset(ctx context.Context, filter struct {
	ID string `json:"id"`
}) Response {
	asset, err := a.wallet.AssetReg.GetAsset(filter.ID)
	if err != nil {
		log.Errorf("getAsset: %v", err)
		return NewErrorResponse(err)
	}

	return NewSuccessResponse(asset)
}

// POST /list-assets
func (a *API) listAssets(ctx context.Context, filter struct {
	Address string `json:"address"`
}) Response {
	assets := make(map[string]uint64)
	var outputs []db.TxOutputs

	err := db.Engine.Select("asset_id, sum(amount) as amount").Where("address = ?", filter.Address).GroupBy("asset_id").Find(&outputs)
	if err != nil {
		log.Errorf("listAssets: %v", err)
		return NewErrorResponse(err)
	}

	for i := range outputs {
		assets[outputs[i].AssetId] = outputs[i].Amount
	}
	return NewSuccessResponse(assets)
}

// POST /list-balances
func (a *API) listBalances(ctx context.Context) Response {
	balances, err := a.wallet.GetAccountBalances("")
	if err != nil {
		return NewErrorResponse(err)
	}
	return NewSuccessResponse(balances)
}

// POST /get-transaction
func (a *API) getTransaction(ctx context.Context, txInfo struct {
	TxID string `json:"tx_id"`
}) Response {
	transaction, err := a.wallet.GetTransactionByTxID(txInfo.TxID)
	if err != nil {
		log.Errorf("getTransaction error: %v", err)
		return NewErrorResponse(err)
	}

	return NewSuccessResponse(transaction)
}

// TxResponse for list-transactions response
type TxResponse struct {
	ID                     string          `json:"id"`
	Timestamp              uint64          `json:"timestamp"`
	BlockID                string          `json:"block_id"`
	BlockHeight            uint64          `json:"block_height"`
	BlockTransactionsCount uint32          `json:"block_transaction_count"`
	Confirmation           uint64          `json:"confirmation"`
	StatusFail             bool            `json:"status_fail"`
	Inputs                 []*db.TxInputs  `json:"inputs"`
	Outputs                []*db.TxOutputs `json:"outputs"`
	Op                     string          `json:"op"`
	Fee                    uint64          `json:"fee"`
	Amount                 uint64          `json:"amount"`
}

// POST /list-transactions
func (a *API) listTransactions(ctx context.Context, filter struct {
	Address    string `json:"address"`
	AssetID    string `json:"asset_id"`
	PageNumber int64  `json:"page_number,omitempty"`
	PageSize   int64  `json:"page_size,omitempty"`
}) Response {

	transactions := []*TxResponse{}
	blocks := []*db.Block{}
	blockHashs := []string{}
	blocksMap := map[string]*db.Block{}
	txIDsMap := map[string]string{} // key:tx_id  value:block_hash
	txIDs := []string{}
	var err error
	var bestBlockHeight uint64

	if bestBlockHeight, err = getBestBlockHeight(); err != nil {
		log.Errorf("list-transactions: %v", err)
		return NewErrorResponse(errors.New("list-transaction err"))
	}

	if txIDsMap, txIDs, err = getTxIDAndBlockHash(filter.Address, filter.AssetID); err != nil {
		log.Errorf("list-transactions: %v", err)
		return NewErrorResponse(errors.New("list-transaction err"))
	}

	if len(txIDs) == 0 {
		log.Infof("list-transactions: no transactions with address: %v, asset_id: %v", filter.Address, filter.AssetID)
		return NewSuccessResponse(transactions)
	}

	// 分页
	sort.Strings(txIDs)
	total := int64(len(txIDs))
	start, end, err := getPagination(total, filter.PageNumber, filter.PageSize)
	if err != nil {
		log.Errorf("list-transactions: %v", err)
		return NewErrorResponse(err)
	}
	returnTxIDs := txIDs[start:end]

	// 获取相关的block
	for i := range returnTxIDs {
		blocksMap[txIDsMap[returnTxIDs[i]]] = nil
	}
	for k := range blocksMap {
		blockHashs = append(blockHashs, k)
	}

	if err = db.Engine.In("hash", blockHashs).Find(&blocks); err != nil {
		log.Errorf("list-transactions: %v", err)
		return NewErrorResponse(errors.New("list-transaction err"))
	}
	for i := range blocks {
		blocksMap[blocks[i].Hash] = blocks[i]
	}

	// 获取结果
	for _, txID := range returnTxIDs {
		var op string
		var inAmount, outAmount uint64
		var inputs []*db.TxInputs
		var outputs []*db.TxOutputs

		// FIXME: 查询需优化
		if err = db.Engine.Where("tx_id = ?", txID).Find(&inputs); err != nil {
			log.Errorf("list-transactions: %v", err)
			continue
		}

		if err = db.Engine.Where("tx_id = ?", txID).Find(&outputs); err != nil {
			log.Errorf("list-transactions: %v", err)
			continue
		}

		for i := range inputs {
			inAmount += inputs[i].Amount
			if inputs[i].Address == filter.Address {
				op = "send"
			}
		}

		for i := range outputs {
			outAmount += outputs[i].Amount
			// 排除找零地址
			if outputs[i].Address == filter.Address && op == "" {
				op = "receive"
			}
		}

		tx := db.Transactions{}
		has, err := db.Engine.Where("tx_id = ?", txID).Get(&tx)
		if err != nil || !has {
			log.Errorf("list-transactions: %v", err)
			continue
		}

		blockHash := txIDsMap[txID]
		txResp := &TxResponse{
			ID:                     txID,
			Timestamp:              blocksMap[blockHash].Timestamp,
			BlockID:                blockHash,
			BlockHeight:            blocksMap[blockHash].Height,
			BlockTransactionsCount: uint32(blocksMap[blockHash].TxCount),
			Confirmation:           bestBlockHeight - blocksMap[blockHash].Height,
			StatusFail:             false,
			Inputs:                 inputs,
			Outputs:                outputs,
			Op:                     op,
			Fee:                    inAmount - outAmount, // 手续费
			Amount:                 tx.Amount,
		}
		transactions = append(transactions, txResp)
	}

	return NewSuccessResponse(transactions)
}

func getPagination(total, pageNumber, pageSize int64) (int64, int64, error) {
	if pageNumber < 0 || pageSize < 0 {
		return 0, 0, errors.New("page params out of range")
	}

	if pageNumber == 0 {
		pageNumber = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	start := (pageNumber - 1) * pageSize
	end := start + pageSize
	if start >= total {
		return 0, 0, errors.New("page params out of range")
	}
	if end >= total {
		end = total - 1
	}

	return start, end, nil
}

func getTxIDAndBlockHash(address, assetID string) (map[string]string, []string, error) {
	txIDsMap := map[string]string{}
	txIDs := []string{}
	var err error

	sql := `
		select tx_id, block_hash from tx_inputs where address = ? and asset_id = ? 
		union
		select tx_id, block_hash from tx_outputs where address = ? and asset_id = ?
	`
	result, err := db.Engine.Query(sql, address, assetID, address, assetID)
	if err != nil {
		return nil, nil, err
	}

	for _, row := range result {
		txID := string(row["tx_id"])
		blockHash := string(row["block_hash"])
		if _, ok := txIDsMap[txID]; !ok {
			txIDsMap[txID] = blockHash
			txIDs = append(txIDs, txID)
		}
	}

	return txIDsMap, txIDs, nil
}

func getBestBlockHeight() (uint64, error) {
	var heightBlcok db.Block
	if _, err := db.Engine.Select("max(height) as height").Get(&heightBlcok); err != nil {
		log.Errorf("list-transactions: %v", err)
		return 0, err
	}

	return heightBlcok.Height, nil
}

// POST /get-unconfirmed-transaction
func (a *API) getUnconfirmedTx(ctx context.Context, filter struct {
	TxID chainjson.HexBytes `json:"tx_id"`
}) Response {
	var tmpTxID [32]byte
	copy(tmpTxID[:], filter.TxID[:])

	txHash := bc.NewHash(tmpTxID)
	txPool := a.chain.GetTxPool()
	txDesc, err := txPool.GetTransaction(&txHash)
	if err != nil {
		return NewErrorResponse(err)
	}

	tx := &BlockTx{
		ID:         txDesc.Tx.ID,
		Version:    txDesc.Tx.Version,
		Size:       txDesc.Tx.SerializedSize,
		TimeRange:  txDesc.Tx.TimeRange,
		Inputs:     []*query.AnnotatedInput{},
		Outputs:    []*query.AnnotatedOutput{},
		StatusFail: false,
	}

	for i := range txDesc.Tx.Inputs {
		tx.Inputs = append(tx.Inputs, a.wallet.BuildAnnotatedInput(txDesc.Tx, uint32(i)))
	}
	for i := range txDesc.Tx.Outputs {
		tx.Outputs = append(tx.Outputs, a.wallet.BuildAnnotatedOutput(txDesc.Tx, i))
	}

	return NewSuccessResponse(tx)
}

type unconfirmedTxsResp struct {
	Total uint64    `json:"total"`
	TxIDs []bc.Hash `json:"tx_ids"`
}

// POST /list-unconfirmed-transactions
func (a *API) listUnconfirmedTxs(ctx context.Context) Response {
	txIDs := []bc.Hash{}

	txPool := a.chain.GetTxPool()
	txs := txPool.GetTransactions()
	for _, txDesc := range txs {
		txIDs = append(txIDs, bc.Hash(txDesc.Tx.ID))
	}

	return NewSuccessResponse(&unconfirmedTxsResp{
		Total: uint64(len(txIDs)),
		TxIDs: txIDs,
	})
}

// RawTx is the tx struct for getRawTransaction
type RawTx struct {
	Version   uint64                   `json:"version"`
	Size      uint64                   `json:"size"`
	TimeRange uint64                   `json:"time_range"`
	Inputs    []*query.AnnotatedInput  `json:"inputs"`
	Outputs   []*query.AnnotatedOutput `json:"outputs"`
	Fee       int64                    `json:"fee"`
}

// POST /decode-raw-transaction
func (a *API) decodeRawTransaction(ctx context.Context, ins struct {
	Tx types.Tx `json:"raw_transaction"`
}) Response {
	tx := &RawTx{
		Version:   ins.Tx.Version,
		Size:      ins.Tx.SerializedSize,
		TimeRange: ins.Tx.TimeRange,
		Inputs:    []*query.AnnotatedInput{},
		Outputs:   []*query.AnnotatedOutput{},
	}

	for i := range ins.Tx.Inputs {
		tx.Inputs = append(tx.Inputs, a.wallet.BuildAnnotatedInput(&ins.Tx, uint32(i)))
	}
	for i := range ins.Tx.Outputs {
		tx.Outputs = append(tx.Outputs, a.wallet.BuildAnnotatedOutput(&ins.Tx, i))
	}

	totalInputBtm := uint64(0)
	totalOutputBtm := uint64(0)
	for _, input := range tx.Inputs {
		if input.AssetID.String() == consensus.BTMAssetID.String() {
			totalInputBtm += input.Amount
		}
	}

	for _, output := range tx.Outputs {
		if output.AssetID.String() == consensus.BTMAssetID.String() {
			totalOutputBtm += output.Amount
		}
	}

	tx.Fee = int64(totalInputBtm) - int64(totalOutputBtm)
	return NewSuccessResponse(tx)
}

// POST /list-unspent-outputs
func (a *API) listUnspentOutputs(ctx context.Context, filter struct {
	ID string `json:"id"`
}) Response {
	accountUTXOs := a.wallet.GetAccountUTXOs(filter.ID)

	UTXOs := []query.AnnotatedUTXO{}
	for _, utxo := range accountUTXOs {
		UTXOs = append([]query.AnnotatedUTXO{{
			AccountID:           utxo.AccountID,
			OutputID:            utxo.OutputID.String(),
			SourceID:            utxo.SourceID.String(),
			AssetID:             utxo.AssetID.String(),
			Amount:              utxo.Amount,
			SourcePos:           utxo.SourcePos,
			Program:             fmt.Sprintf("%x", utxo.ControlProgram),
			ControlProgramIndex: utxo.ControlProgramIndex,
			Address:             utxo.Address,
			ValidHeight:         utxo.ValidHeight,
			Alias:               a.wallet.AccountMgr.GetAliasByID(utxo.AccountID),
			AssetAlias:          a.wallet.AssetReg.GetAliasByID(utxo.AssetID.String()),
			Change:              utxo.Change,
		}}, UTXOs...)
	}

	return NewSuccessResponse(UTXOs)
}

// return gasRate
func (a *API) gasRate() Response {
	gasrate := map[string]int64{"gas_rate": consensus.VMGasRate}
	return NewSuccessResponse(gasrate)
}
