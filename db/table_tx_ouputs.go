package db

type (
	TxOutputs struct {
		TxId      string `xorm:"tx_id"`
		BlockHash string `xorm:"block_hash"`
		Address   string `xorm:"address"`
		AssetId   string `xorm:"asset_id"`
		Amount    uint64 `xorm:"amount"`
		OutputId  string `xorm:"output_id"`
		Type      uint64 `xorm:"type"`
	}
)

func (t *TxOutputs) TableName() string {
	return "tx_outputs"
}

func (db *DB) SaveTxOutputs(t *TxOutputs) (err error) {
	_, err = db.engine.Insert(t)
	if err != nil {
		return
	}
	return
}
