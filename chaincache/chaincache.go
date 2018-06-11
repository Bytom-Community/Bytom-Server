package chaincache

import (
	"context"
	"sync"

	"github.com/bytom/account"

	"github.com/bytom/asset"
	"github.com/bytom/blockchain/query"
	"github.com/bytom/database/leveldb"
	"github.com/bytom/protocol"
	"github.com/bytom/protocol/bc"
	"github.com/bytom/protocol/bc/types"
	w "github.com/bytom/wallet"

	"time"

	log "github.com/sirupsen/logrus"
	cmn "github.com/tendermint/tmlibs/common"
)

type ChainCache struct {
	sync.RWMutex

	store               *leveldb.Store
	chain               *protocol.Chain
	wallet              *w.Wallet
	BlockChain          []*types.Block
	TransactionsOutputs map[bc.Hash][]*query.AnnotatedOutput
	exitCh              chan bool
}

func NewChainCache(store *leveldb.Store, chain *protocol.Chain, wallet *w.Wallet) *ChainCache {
	c := &ChainCache{
		store:  store,
		chain:  chain,
		wallet: wallet,
		exitCh: make(chan bool),
	}
	c.Lock()
	c.updateChain()
	c.Unlock()

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
			c.Lock()
			c.updateChain()
			c.Unlock()
		}
	}
exit:
	log.Info("CahinCache goroutine exiting")
	workTicker.Stop()
}

func (c *ChainCache) updateChain() {
	bc, txOutputs := c.readChain()
	c.BlockChain = bc
	c.TransactionsOutputs = txOutputs
}

func (c *ChainCache) readChain() (blocks []*types.Block, txOutputs map[bc.Hash][]*query.AnnotatedOutput) {

	bcHash := c.chain.BestBlockHash()
	height := c.chain.BestBlockHeight()
	txOutputs = make(map[bc.Hash][]*query.AnnotatedOutput)

	for i := height; i > 0; i-- {
		block, err := c.store.GetBlock(bcHash)
		if err != nil {
			cmn.Exit(cmn.Fmt("Failed to start switch: %v", err))
		}
		blocks = append(blocks, block)

		var outpus = []*query.AnnotatedOutput{}
		for _, tx := range block.Transactions {
			for i := range tx.Outputs {
				outpus = append(outpus, c.wallet.BuildAnnotatedOutput(tx, i))
			}
		}
		txOutputs[block.BlockHeader.Hash()] = outpus

		bcHash = &block.PreviousBlockHash
	}
	return blocks, txOutputs
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

func (c *ChainCache) ListTransactions(address, assetID string) []*types.Tx {
	txs := []*types.Tx{}
	c.RLock()
	defer c.RUnlock()
	for _, block := range c.BlockChain {
		for _, tx := range block.Transactions {
			var outpus = []*query.AnnotatedOutput{}
			for i := range tx.Outputs {
				outpus = append(outpus, c.wallet.BuildAnnotatedOutput(tx, i))
			}

			for _, v := range outpus {
				if v.Address == address && v.AssetID.String() == assetID {
					txs = append(txs, tx)
				}
			}
		}
	}

	return txs
}

func (c *ChainCache) ListTransaction(txID string) map[string]interface{} {
	var TX = make(map[string]interface{})
	c.RLock()
	defer c.RUnlock()
exit:
	for _, block := range c.BlockChain {
		for _, tx := range block.Transactions {
			var inputs = []*query.AnnotatedInput{}
			var outputs = []*query.AnnotatedOutput{}

			if tx.ID.String() == txID {
				for i := range tx.Inputs {
					inputs = append(inputs, c.wallet.BuildAnnotatedInput(tx, uint32(i)))
				}

				for i := range tx.Outputs {
					outputs = append(outputs, c.wallet.BuildAnnotatedOutput(tx, i))
				}
				TX["block"] = block
				TX["inputs"] = inputs
				TX["outputs"] = outputs
				break exit
			}
		}
	}

	return TX
}

func (c *ChainCache) Close() {
	close(c.exitCh)
	c.BlockChain = []*types.Block{}
	c.TransactionsOutputs = make(map[bc.Hash][]*query.AnnotatedOutput)
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
