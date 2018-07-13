package db

import (
	"fmt"
	"errors"
	"strings"
)

const BTMAssetID = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

var mydb *DB
var outputs []TxOutputs
var inputs []TxInputs
var block []Block

type TX struct {
	ID                     string    `json:"ID,omitempty"`
	Timestamp              uint64    `json:"timestamp,omitempty"`
	BlockID                string    `json:"blockID,omitempty"`
	BlockHeight            uint64    `json:"blockHeight,omitempty"`
	Position               uint32    `json:"position,omitempty"`
	BlockTransactionsCount uint32    `json:"blockTransactionsCount,omitempty"`
	Confirmation           uint64    `json:"confirmation,omitempty"`
	StatusFail             bool      `json:"statusFail,omitempty"`
	Inputs                 []*Input  `json:"inputs,omitempty"`
	Outputs                []*Output `json:"outputs,omitempty"`
	Op                     string    `json:"op,omitempty"`
	Fee                    uint64    `json:"fee,omitempty"`
}

// transactions
type Input struct {
	Type          string `json:"type,omitempty"`
	AssetID       string `json:"assetID,omitempty"`
	Amount        uint64 `json:"amount,omitempty"`
	Address       string `json:"address,omitempty"`
	SpentOutputID string `json:"spentOutputID,omitempty"`
}
type Output struct {
	Type     string `json:"type,omitempty"`
	AssetID  string `json:"assetID,omitempty"`
	Amount   uint64 ` json:"amount,omitempty"`
	Address  string `json:"address,omitempty"`
	OutputID string `json:"OutputID,omitempty"`
}

func (mydb *DB) GetTransactions(address, AssetID string) ([]*TX, error) {
	//Get tx_id and block_hash from DB.tx_inputs, put them in map TxIn
		TxIn, err := mydb.GetTxIn(address, AssetID)
	//Get tx_id and block_hash from DB.tx_outputs, put them in map TxOut
	TxOut, err := mydb.GetTxOut(address, AssetID)
	//merge TxIn and TxOut to Tx
	for k, v := range TxOut {
		TxIn[k] = v
	}
	Tx := TxIn
	//Get  TX'sTransactions  from DB
	Transactions, err := mydb.GetTransaction(Tx, address, AssetID)
	if err != nil {
		fmt.Println(err)
	}
	return Transactions, err
}
func (mydb *DB) GetBestBlockHeight() (uint64, error) {
	//Get BestBlockHeight from DB.block table from DB.tx_outputs
	err := mydb.Engine.Select("max(height) as height").Find(&block)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if len(block) != 1 {
		return 0, errors.New("The max block height value is not one value. There are  " + string(len(block)) + "  value.")
	}
	return block[0].Height, nil
}

// GetTxOut(map tx_id and block_hash) by address and AssetID
func (mydb *DB) GetTxOut(address string, assetID string) (map[string]string, error) {
	TxOut := make(map[string]string)
	// select tx_id, block_hash from tx_outputs where address = ? and asset_id=?;
	err := mydb.Engine.Select("tx_id , block_hash").Where("address = ?", address).And("asset_id=?", assetID).Find(&outputs)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for i := range outputs {
		TxOut[outputs[i].TxId] = outputs[i].BlockHash
	}
	return TxOut, nil
}

// GetTxIn(map tx_id and block_hash) by address and AssetID from DB.tx_inputs,
func (mydb *DB) GetTxIn(address string, assetID string) (map[string]string, error) {
	TxIn := make(map[string]string)
	// select tx_id, block_hash from tx_inputs where address = ? and asset_id=?;
	err := mydb.Engine.Select("tx_id , block_hash").Where("address = ?", address).And("asset_id=?", assetID).Find(&inputs)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for i := range inputs {
		TxIn[inputs[i].TxId] = inputs[i].BlockHash
	}
	return TxIn, nil
}
func (mydb *DB) GetTransaction(Tx map[string]string, address, AssetID string) ([]*TX, error) {
	Transactions := []*TX{}
	BestBlockHeight, err := mydb.GetBestBlockHeight()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var op string
	for TxID, block_hash := range Tx {
		//initial input,output,block nil,op ""
		op = ""
		inputs = nil
		outputs = nil
		block = nil
		// select inputs from tx_inputs where tx_id=?;
		err := mydb.Engine.Select("*").Where("tx_id = ?", TxID).Find(&inputs)
		if err != nil {
			return nil, err
		}
		NewInputs := []*Input{}
		Amount := uint64(0)
		for i := range inputs {
			input := &Input{
				Address:       inputs[i].Address,
				AssetID:       inputs[i].AssetId,
				Amount:        inputs[i].Amount,
				SpentOutputID: inputs[i].SpentOutputId,
				Type:          inputs[i].Type,
			}
			if strings.EqualFold(inputs[i].AssetId, BTMAssetID) {
				Amount += inputs[i].Amount
			}
			if input.Address == address {
				op = "send"
			}
			NewInputs = append(NewInputs, input)
		}

		// select outputs from tx_outputs where tx_id=?;
		err = mydb.Engine.Select("*").Where("tx_id = ?", TxID).Find(&outputs)
		if err != nil {
			return nil, err
		}
		NewOutputs := []*Output{}
		for i := range outputs {
			output := &Output{
				Address:  outputs[i].Address,
				AssetID:  outputs[i].AssetId,
				Amount:   outputs[i].Amount,
				OutputID: outputs[i].OutputId,
				Type:     outputs[i].Type,
			}
			if strings.EqualFold(outputs[i].AssetId, BTMAssetID) {
				Amount -= outputs[i].Amount
			}
			if output.Address == address && op == "" {
				op = "receive"
			}
			NewOutputs = append(NewOutputs, output)
		}
		// Get  transaction from block where block_hash=?;
		Txblock := new(Block)
		result, err := mydb.Engine.Where("hash = ?", block_hash).Get(Txblock)
		if err != nil {
			return nil, err
		}
		//block should be unique one
		if result == false {
			return nil, errors.New("Can not get  block from database with block_hash:" + block_hash)
		}
		BlockTx := &TX{
			ID:                     TxID,
			Timestamp:              Txblock.Timestamp,
			BlockHeight:            Txblock.Height,
			BlockID:                Txblock.Hash,
			BlockTransactionsCount: uint32(Txblock.TxCount),
			Confirmation:           BestBlockHeight - Txblock.Height,
			Op:                     op,
			Fee:                    Amount,
			Inputs:                 NewInputs,
			Outputs:                NewOutputs,
		}
		Transactions = append(Transactions, BlockTx)
	}
	return Transactions, nil
}
