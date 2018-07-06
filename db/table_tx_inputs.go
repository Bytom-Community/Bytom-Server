package db

type (
	TxInputs struct {
		TxId          string `xorm:"tx_id"`
		BlockHash     string `xorm:"block_hash"`
		Address       string `xorm:"address"`
		AssetId       string `xorm:"asset_id"`
		Amount        uint64 `xorm:"amount"`
		SpentOutputId string `xorm:"spent_output_id"`
		Type          uint64 `xorm:"type"`
	}
)

func (t *TxInputs) TableName() string {
	return "tx_inputs"
}

func (db *DB) SaveTxInputs(t *TxInputs) (err error) {
	_, err = db.engine.Insert(t)
	if err != nil {
		return
	}
	return
}
