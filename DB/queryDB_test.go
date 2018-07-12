package db

import (
	"testing"
	"fmt"
)

func TestGetTransactions(t *testing.T) {
	mydb, err := mockDB()
	if err != nil {
		t.Fatal(err)
	}
	Tx, err := mydb.GetTransactions("bm1q5p9d4gelfm4cc3zq3slj7vh2njx23ma2cf866j", "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	if err != nil {
		t.Fatal(err)
	}
	showTransaction(Tx)
}

func showTransaction(Tx []*TX) {
	for i := range Tx {
		fmt.Println("Tx:", i, " ID:", Tx[i].ID)
		fmt.Println("        blockHeight:  ", Tx[i].BlockHeight)
		fmt.Println("        blockID:     ", Tx[i].BlockID)
		fmt.Println("        TxCount:    ", Tx[i].BlockTransactionsCount)
		fmt.Println("        Confirmation ", Tx[i].Confirmation)
		fmt.Println("        fee:     ", Tx[i].Fee)
		fmt.Println("        op       ", Tx[i].Op)
		for j := range Tx[i].Inputs {
			fmt.Println("           TxInput:", j)
			fmt.Println("                             address:", Tx[i].Inputs[j].Address)
			fmt.Println("                             amount:", Tx[i].Inputs[j].Amount)
			fmt.Println("                             assetID:", Tx[i].Inputs[j].AssetID)
			fmt.Println("                             spentOutputID:", Tx[i].Inputs[j].SpentOutputID)
			fmt.Println("                             type:", Tx[i].Inputs[j].Type)
		}
		for k := range Tx[i].Outputs {
			fmt.Println("           TxOutput:", k)
			fmt.Println("                             OutputID:", Tx[i].Outputs[k].OutputID)
			fmt.Println("                             address:", Tx[i].Outputs[k].Address)
			fmt.Println("                             amount:", Tx[i].Outputs[k].Amount)
			fmt.Println("                             assetID:", Tx[i].Outputs[k].AssetID)
			//fmt.Println("                             Position:",Tx[i].Outputs[k].position)
			fmt.Println("                             type:", Tx[i].Outputs[k].Type)
		}
	}
}

func mockDB() (*DB, error) {
	mydb, err := NewDB("localhost", "root", "test1234", "3306", "bytom-server")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//连接测试
	if err := mydb.Engine.Ping(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return mydb, nil
}
