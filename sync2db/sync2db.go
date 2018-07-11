package sync2db

import (
	"strings"
	"time"

	"github.com/bytom/database/leveldb"
	"github.com/bytom/db"
	"github.com/bytom/protocol"
	w "github.com/bytom/wallet"

	log "github.com/sirupsen/logrus"
	cmn "github.com/tendermint/tmlibs/common"
)

type Sync2DB struct {
	store  *leveldb.Store
	chain  *protocol.Chain
	wallet *w.Wallet
	exitCh chan bool
}

func NewSync2DB(store *leveldb.Store, chain *protocol.Chain, wallet *w.Wallet) *Sync2DB {
	return &Sync2DB{
		store:  store,
		chain:  chain,
		wallet: wallet,
		exitCh: make(chan bool),
	}
}

func (s *Sync2DB) Run() {
	log.Info("Sync2DB goroutine running")
	s.runSync()
	workTicker := time.NewTicker(time.Minute)
	for {
		select {
		case <-s.exitCh:
			goto exit
		case <-workTicker.C:
			s.runSync()
		}
	}
exit:
	log.Info("Sync2DB goroutine exiting")
	workTicker.Stop()
}

func (s *Sync2DB) runSync() {
	bcHash := s.chain.BestBlockHash()
	height := s.chain.BestBlockHeight()
	existCount := 0
	log.Infof("Sync2DB running sync to db， current height:%v", height)

	for i := height; i > 0; i-- {
		block, err := s.store.GetBlock(bcHash)
		if err != nil {
			cmn.Exit(cmn.Fmt("Failed to get block from store: %v", err))
		}
		if existCount >= 3 {
			return
		}
		blockID, err := block.Hash().MarshalText()
		if err != nil {
			cmn.Exit(cmn.Fmt("Failed to MarshalText from hash :%v", err))
		}
		exist, err := db.Engine.Exist(&db.Block{
			Hash: string(blockID),
		})
		if err != nil {
			cmn.Exit(cmn.Fmt("Failed to get block exist from db: %v", err))
		}
		log.Info("current sync2db block id:", string(blockID))
		if !exist {
			var txIds []string
			for _, v := range block.Transactions {
				txIds = append(txIds, v.ID.String())
			}
			_, err = db.Engine.Insert(&db.Block{
				Hash:              string(blockID),
				Version:           block.Version,
				Height:            block.Height,
				PreviousBlockHash: block.PreviousBlockHash.String(),
				Timestamp:         block.Timestamp,
				Nonce:             block.Nonce,
				Bits:              block.Bits,
				TxCount:           len(block.Transactions),
				TxIds:             strings.Join(txIds, ","),
			})
			if err != nil {
				cmn.Exit(cmn.Fmt("Failed to insert block from db: %v", err))
			}
		} else {
			existCount++
		}

		for _, tx := range block.Transactions {
			var txid = tx.ID.String()

			// inputs
			for i := range tx.Inputs {
				input := s.wallet.BuildAnnotatedInput(tx, uint32(i))
				has, err := db.Engine.Exist(&db.TxInputs{
					TxId:          txid,
					BlockHash:     string(blockID),
					Address:       input.Address,
					AssetId:       input.AssetID.String(),
					SpentOutputId: input.SpentOutputID.String(),
					Type:          input.Type,
				})
				if err != nil {
					cmn.Exit(cmn.Fmt("Failed to exist block from db: %v", err))
				}
				assetDefinition, err := input.AssetDefinition.MarshalJSON()
				if err != nil {
					cmn.Exit(cmn.Fmt("Failed to get AssetDefinition from json: %v", err))
				}
				if !has {
					_, err = db.Engine.Insert(&db.TxInputs{
						TxId:            txid,
						BlockHash:       string(blockID),
						Address:         input.Address,
						AssetId:         input.AssetID.String(),
						Amount:          input.Amount,
						AssetDefinition: string(assetDefinition),
						SpentOutputId:   input.SpentOutputID.String(),
						Type:            input.Type,
					})
					if err != nil {
						cmn.Exit(cmn.Fmt("Failed to insert inputs from db: %v", err))
					}
				}
			}

			// outputs
			for i := range tx.Outputs {
				output := s.wallet.BuildAnnotatedOutput(tx, i)
				has, err := db.Engine.Exist(&db.TxOutputs{
					TxId:      txid,
					BlockHash: string(blockID),
					Address:   output.Address,
					AssetId:   output.AssetID.String(),
					OutputId:  output.OutputID.String(),
					Type:      output.Type,
				})
				if err != nil {
					cmn.Exit(cmn.Fmt("Failed to exist block from db: %v", err))
				}
				assetDefinition, err := output.AssetDefinition.MarshalJSON()
				if err != nil {
					cmn.Exit(cmn.Fmt("Failed to get AssetDefinition from json: %v", err))
				}
				if !has {
					_, err = db.Engine.Insert(&db.TxOutputs{
						TxId:            txid,
						BlockHash:       string(blockID),
						Address:         output.Address,
						AssetId:         output.AssetID.String(),
						Amount:          output.Amount,
						AssetDefinition: string(assetDefinition),
						OutputId:        output.OutputID.String(),
						Type:            output.Type,
					})
					if err != nil {
						cmn.Exit(cmn.Fmt("Failed to insert outputs from db: %v", err))
					}
				}
			}
		}
		bcHash = &block.PreviousBlockHash
	}
}

func (s *Sync2DB) Close() {
	close(s.exitCh)
}
