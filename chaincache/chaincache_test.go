package chaincache

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/bytom/database/leveldb"
	"github.com/bytom/protocol"
	w "github.com/bytom/wallet"
	"testing"

	"github.com/Bytom-Community/Bytom-Mobile-Wallet/util"
	//"github.com/bytom/protocol/bc/types"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
)

var (
	levelDBDir   = "/Users/coral/Library/Bytom/data"
	stime, etime int64
	store        *leveldb.Store
	chain        *protocol.Chain
	wallet       *w.Wallet

	c *ChainCache
)

func init() {
	var err error
	coreDB := dbm.NewDB("core", "leveldb", levelDBDir)
	store = leveldb.NewStore(coreDB)

	txPool := protocol.NewTxPool()
	chain, err = protocol.NewChain(store, txPool)
	if err != nil {
		cmn.Exit(cmn.Fmt("Failed to create chain structure: %v", err))
	}

	wallet = new(w.Wallet)

	c = NewChainCache(store, chain, wallet)
}

func TestListAssets(t *testing.T) {
	stime = time.Now().UnixNano()
	d := c.ListAssets("bm1q5p9d4gelfm4cc3zq3slj7vh2njx23ma2cf866j")
	etime = time.Now().UnixNano()
	totalTime := (etime - stime) / 1e6
	if len(d) != 0 {
		for k, v := range d {
			t.Logf("assetID:%v, amount:%v \n", k, v)
		}
		t.Logf("run time: %v ms", totalTime)
	}
}

func TestListTransactions(t *testing.T) {
	stime = time.Now().UnixNano()
	d := c.ListTransactions("bm1q5p9d4gelfm4cc3zq3slj7vh2njx23ma2cf866j", "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	etime = time.Now().UnixNano()
	totalTime := (etime - stime) / 1e6
	if len(d) != 0 {
		b := util.JsonEncode(d)
		t.Log(string(b))
		t.Logf("run time: %v ms", totalTime)
	}
}

func TestGetTransactionByID(t *testing.T) {
	stime = time.Now().UnixNano()
	d := c.getTransactionByID("0ebb5cfd1bc9c9bb2417edf69f61cb68abcd33318293ff4ae18dc7b483f6b941")
	etime = time.Now().UnixNano()
	totalTime := (etime - stime) / 1e6
	if d != nil {
		t.Log(d.BlockHeight)
		t.Log(d.BlockID.String())
		t.Logf("run time: %v ms", totalTime)
	}
}

func TestCount(t *testing.T) {
	address := "bm1q5p9d4gelfm4cc3zq3slj7vh2njx23ma2cf866j"
	assetID := "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	inputCount, outputCount := 0, 0

	for txid, inputs := range c.TransactionsInput {
		for _, tx := range inputs {
			if tx.Address == address && tx.AssetID.String() == assetID {
				t.Log("input: ", txid)
				inputCount++
			}
		}
	}

	for txid, outputs := range c.TransactionsOutputs {
		for _, tx := range outputs {
			if tx.Address == address && tx.AssetID.String() == assetID {
				t.Log("output: ", txid)
				outputCount++
			}
		}
	}
	t.Log("inputCount:", inputCount)
	t.Log("outputCount:", outputCount)

	t.Log("BlockChain:", len(c.BlockChain))
	t.Log("TransactionsInput:", len(c.TransactionsInput))
	t.Log("TransactionsOutputs:", len(c.TransactionsOutputs))
}

//
//func TestListTransaction(t *testing.T) {
//	//height := c.chain.BestBlockHeight()
//	stime = time.Now().UnixNano()
//	d := c.ListTransaction("795e78a66ab73c209e13635a057810098455bf15b68107f2df93edf910185bf6")
//	etime = time.Now().UnixNano()
//	totalTime := (etime - stime) / 1e6
//	if len(d) != 0 {
//		block, ok := d["block"].(*types.Block)
//		if !ok {
//			panic("can not cover block to *types.Block")
//		}
//		t.Logf("height:%v, timestamp:%v, len:%v, confirmations:%v\n",
//			block.Height, block.Timestamp, len(block.Transactions), chain.BestBlockHeight()-block.Height)
//
//		var TX = make(map[string]interface{})
//		TX["inputs"] = d["inputs"]
//		TX["outputs"] = d["outputs"]
//		printJSON(TX)
//		t.Logf("run time: %v ms", totalTime)
//	}
//}

func printJSON(data interface{}) {
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		panic("invalid type assertion")
		os.Exit(1)
	}

	rawData, err := json.MarshalIndent(dataMap, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(rawData))
}
