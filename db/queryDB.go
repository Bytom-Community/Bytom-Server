package db

import (
	"fmt"
	"errors"
	"strings"
	"github.com/Bytom-Server/rpc/pb"
)

const BTMAssetID = "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

var mydb *DB
var outputs []TxOutputs
var inputs []TxInputs
var block []Block

func (mydb *DB) GetTransactions(address, AssetID string) ([]*rpcpb.TX, error) {
	TxIn := make(map[string]string)
	TxIn, err := mydb.GetTxIn(address, AssetID)

	TxOut := make(map[string]string)
	TxOut, err = mydb.GetTxOut(address, AssetID)
	//merge TxIn and TxOut
	for k, v := range TxOut {
		TxIn[k] = v
	}
	Tx := TxIn
	Transactions, err := mydb.GetTransaction(Tx)
	if err != nil {
		fmt.Println(err)
	}
	return Transactions, err
}
func (mydb *DB) GetBestBlockHeight() (uint64, error) {
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

// GetTxOut by address and AssetID
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

// GetTxIn by address and AssetID
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
func (mydb *DB) GetTransaction(Tx map[string]string) ([]*rpcpb.TX, error) {
	Transactions := []*rpcpb.TX{}
	BestBlockHeight, err := mydb.GetBestBlockHeight()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for TxID, block_hash := range Tx {
		//initial input,output,block nil
		inputs = nil
		outputs = nil
		block = nil
		// select inputs from tx_inputs where tx_id=?;
		err := mydb.Engine.Select("*").Where("tx_id = ?", TxID).Find(&inputs)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		NewInputs := []*rpcpb.Input{}
		Amount := uint64(0)
		for i := range inputs {
			input := &rpcpb.Input{
				Address:       inputs[i].Address,
				AssetID:       inputs[i].AssetId,
				Amount:        inputs[i].Amount,
				SpentOutputID: inputs[i].SpentOutputId,
				Type:          inputs[i].Type,
			}
			if strings.EqualFold(inputs[i].AssetId, BTMAssetID) {
				Amount += inputs[i].Amount
			}
			NewInputs = append(NewInputs, input)
		}

		// select outputs from tx_outputs where tx_id=?;
		err = mydb.Engine.Select("*").Where("tx_id = ?", TxID).Find(&outputs)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		NewOutputs := []*rpcpb.Output{}
		for i := range outputs {
			output := &rpcpb.Output{
				Address:  outputs[i].Address,
				AssetID:  outputs[i].AssetId,
				Amount:   outputs[i].Amount,
				OutputID: outputs[i].OutputId,
				Type:     outputs[i].Type,
			}
			if strings.EqualFold(outputs[i].AssetId, BTMAssetID) {
				Amount -= outputs[i].Amount
			}
			NewOutputs = append(NewOutputs, output)
		}
		// select  transaction from block where block_hash=?;
		err = mydb.Engine.Select("*").Where("hash = ?", block_hash).Find(&block)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		//block should be unique one
		if len(block) != 1 {
			return nil, errors.New("Error happened when get  block from block_hash.The  return block value is not one value. There are  " + string(len(block)) + "  value with block_hash:" + block_hash)
		}
		Txblock := block[0]
		BlockTx := &rpcpb.TX{
			ID:                     Txblock.TxIds,
			Timestamp:              Txblock.Timestamp,
			BlockHeight:            Txblock.Height,
			BlockID:                Txblock.Hash,
			BlockTransactionsCount: uint32(Txblock.TxCount),
			Confirmation:           BestBlockHeight - Txblock.Height,
			Fee:                    Amount,
			Inputs:                 NewInputs,
			Outputs:                NewOutputs,
		}
		Transactions = append(Transactions, BlockTx)
	}
	return Transactions, nil
}
