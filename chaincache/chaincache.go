package chaincache

import (
	"context"
	"sync"
	"time"

	"github.com/bytom/account"
	"github.com/bytom/asset"
	"github.com/bytom/blockchain/query"
	"github.com/bytom/database/leveldb"
	"github.com/bytom/protocol"
	"github.com/bytom/protocol/bc"
	"github.com/bytom/protocol/bc/types"
	w "github.com/bytom/wallet"

	log "github.com/sirupsen/logrus"
	cmn "github.com/tendermint/tmlibs/common"
)

type ChainCache struct {
	sync.RWMutex

	store               *leveldb.Store
	chain               *protocol.Chain
	wallet              *w.Wallet
	BlockChain          map[bc.Hash]*types.Block
	TransactionsInput   map[string][]*query.AnnotatedInput
	TransactionsOutputs map[string][]*query.AnnotatedOutput
	exitCh              chan bool
}

func NewChainCache(store *leveldb.Store, chain *protocol.Chain, wallet *w.Wallet) *ChainCache {
	c := &ChainCache{
		store:  store,
		chain:  chain,
		wallet: wallet,
		exitCh: make(chan bool),
	}
	c.updateChain()

	go c.syncChainRoutine()
	return c
}

func (c *ChainCache) syncChainRoutine() {
	log.Info("CahinCache goroutine running")

	workTicker := time.NewTicker(time.Minute)
	for {
		select {
		case <-c.exitCh:
			goto exit
		case <-workTicker.C:
			c.updateChain()
		}
	}
exit:
	log.Info("CahinCache goroutine exiting")
	workTicker.Stop()
}

func (c *ChainCache) updateChain() {
	bc, txInput, txOutputs := c.readChain()
	c.Lock()
	c.BlockChain = bc
	c.TransactionsInput = txInput
	c.TransactionsOutputs = txOutputs
	c.Unlock()
}

func (c *ChainCache) readChain() (blocks map[bc.Hash]*types.Block, txInput map[string][]*query.AnnotatedInput, txOutputs map[string][]*query.AnnotatedOutput) {

	bcHash := c.chain.BestBlockHash()
	height := c.chain.BestBlockHeight()
	blocks = make(map[bc.Hash]*types.Block)
	txInput = make(map[string][]*query.AnnotatedInput)
	txOutputs = make(map[string][]*query.AnnotatedOutput)

	for i := height; i > 0; i-- {
		block, err := c.store.GetBlock(bcHash)
		if err != nil {
			cmn.Exit(cmn.Fmt("Failed to get block from store: %v", err))
		}
		blocks[block.Hash()] = block

		for _, tx := range block.Transactions {
			var txid = tx.ID.String()
			var input = []*query.AnnotatedInput{}
			var outpus = []*query.AnnotatedOutput{}

			for i := range tx.Inputs {
				input = append(input, c.wallet.BuildAnnotatedInput(tx, uint32(i)))
			}
			for i := range tx.Outputs {
				outpus = append(outpus, c.wallet.BuildAnnotatedOutput(tx, i))
			}
			txInput[txid] = input
			txOutputs[txid] = outpus
		}

		bcHash = &block.PreviousBlockHash
	}
	return blocks, txInput, txOutputs
}

func (c *ChainCache) BestBlockHeight() uint64 {
	return c.chain.BestBlockHeight()
}

func (c *ChainCache) ListAssets(address string) map[string]uint64 {
	assets := make(map[string]uint64)
	c.RLock()
	defer c.RUnlock()
	for _, outputs := range c.TransactionsOutputs {
		for _, tx := range outputs {
			if tx.Address == address {
				assets[tx.AssetID.String()] += tx.Amount
			}
		}
	}
	return assets
}

func (c *ChainCache) ListTransactions(address, assetID string) []*query.AnnotatedTx {
	var transactions = []*query.AnnotatedTx{}
	c.RLock()
	defer c.RUnlock()
	for txid, inputs := range c.TransactionsInput {
		for _, tx := range inputs {
			if tx.Address == address && tx.AssetID.String() == assetID {
				transaction := c.getTransactionByID(txid)
				transactions = append(transactions, transaction)
			}
		}
	}

	for txid, outputs := range c.TransactionsOutputs {
		for _, tx := range outputs {
			if tx.Address == address && tx.AssetID.String() == assetID {
				transaction := c.getTransactionByID(txid)
				transactions = append(transactions, transaction)
			}
		}
	}

	return transactions
}

func (c *ChainCache) getTransactionByID(txID string) *query.AnnotatedTx {
	var transaction = new(query.AnnotatedTx)
exit:
	for blockHash, block := range c.BlockChain {
		for _, tx := range block.Transactions {

			if tx.ID.String() == txID {
				var inputs = []*query.AnnotatedInput{}
				var outputs = []*query.AnnotatedOutput{}

				for i := range tx.Inputs {
					inputs = append(inputs, c.wallet.BuildAnnotatedInput(tx, uint32(i)))
				}
				for i := range tx.Outputs {
					outputs = append(outputs, c.wallet.BuildAnnotatedOutput(tx, i))
				}

				transaction.BlockID = blockHash
				transaction.ID = tx.ID
				transaction.BlockHeight = block.Height
				transaction.Timestamp = block.Timestamp
				transaction.BlockTransactionsCount = uint32(len(block.Transactions))
				transaction.Inputs = inputs
				transaction.Outputs = outputs
				transaction.StatusFail = false
				break exit
			}
		}
	}

	return transaction
}

//
//func (c *ChainCache) ListTransaction(txID string) map[string]interface{} {
//	var TX = make(map[string]interface{})
//	c.RLock()
//	defer c.RUnlock()
//exit:
//	for _, block := range c.BlockChain {
//		for _, tx := range block.Transactions {
//			var inputs = []*query.AnnotatedInput{}
//			var outputs = []*query.AnnotatedOutput{}
//
//			if tx.ID.String() == txID {
//				for i := range tx.Inputs {
//					inputs = append(inputs, c.wallet.BuildAnnotatedInput(tx, uint32(i)))
//				}
//
//				for i := range tx.Outputs {
//					outputs = append(outputs, c.wallet.BuildAnnotatedOutput(tx, i))
//				}
//				TX["block"] = block
//				TX["inputs"] = inputs
//				TX["outputs"] = outputs
//				break exit
//			}
//		}
//	}
//
//	return TX
//}

func (c *ChainCache) Close() {
	close(c.exitCh)
	c.BlockChain = make(map[bc.Hash]*types.Block)
	c.TransactionsInput = make(map[string][]*query.AnnotatedInput)
	c.TransactionsOutputs = make(map[string][]*query.AnnotatedOutput)
}

func (c *ChainCache) FindAssetByAlias(alias string) (*asset.Asset, error) {
	asset, err := c.wallet.AssetReg.FindByAlias(alias)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (c *ChainCache) FindAccountByAlias(ctx context.Context, alias string) (*account.Account, error) {
	acc, err := c.wallet.AccountMgr.FindByAlias(ctx, alias)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (c *ChainCache) GetChain() *protocol.Chain {
	return c.chain
}
