package sync2db

import (
	"testing"

	"github.com/bytom/database/leveldb"
	"github.com/bytom/db"
	"github.com/bytom/protocol"
	w "github.com/bytom/wallet"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"time"
)

var (
	levelDBDir   = "/Users/coral/Library/Bytom/data"
	stime, etime int64
	store        *leveldb.Store
	chain        *protocol.Chain
	wallet       *w.Wallet
	d            *db.DB
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

	d, err = db.NewDB("localhost", "root", "1", "3306", "btmwallet")
	if err != nil {
		cmn.Exit(err.Error())
	}
}

func TestSync2DB_Run(t *testing.T) {
	syc := NewSync2DB(store, chain, wallet, d)

	go func() {
		defer syc.Close()
		syc.Run()
	}()
	time.Sleep(time.Minute)
}
