package db

type (
	TxOutputs struct {
		TxID      string `xorm:"tx_id"`
		BlockHash string `xorm:"block_hash"`
		Address   string `xorm:"address"`
		AssetID   string `xorm:"asset_id"`
		Amount    uint64 `xorm:"amount"`
		OutputID  string `xorm:"output_id"`
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

// GetAssetsByAddress get assets by address
func (db *DB) GetAssetsByAddress(address string) (map[string]uint64, error) {
	// select asset_id, sum(amount) from tx_outputs where address = ? group by asset_id;

	assets := make(map[string]uint64)
	var outputs []TxOutputs
	err := db.engine.Select("asset_id, sum(amount) as amount").Where("address = ?", address).GroupBy("asset_id").Find(&outputs)
	if err != nil {
		return assets, err
	}
	for i := range outputs {
		assets[outputs[i].AssetID] = outputs[i].Amount
	}
	return assets, nil
}
