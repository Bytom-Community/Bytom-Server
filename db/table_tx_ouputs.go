package db

type (
	TxOutputs struct {
		//Id              int    `xorm:"<-"`
		TxId            string `xorm:"tx_id"`
		BlockHash       string `xorm:"block_hash"`
		Address         string `xorm:"address"`
		AssetId         string `xorm:"asset_id"`
		Amount          uint64 `xorm:"amount"`
		AssetDefinition string `xorm:"asset_definition"`
		OutputId        string `xorm:"output_id"`
		Type            string `xorm:"type"`
	}
)

func (t *TxOutputs) TableName() string {
	return "tx_outputs"
}

// GetAssetsByAddress get assets by address
func GetAssetsByAddress(address string) (map[string]uint64, error) {
	// select asset_id, sum(amount) from tx_outputs where address = ? group by asset_id;

	assets := make(map[string]uint64)
	var outputs []TxOutputs
	err := Engine.Select("asset_id, sum(amount) as amount").Where("address = ?", address).GroupBy("asset_id").Find(&outputs)
	if err != nil {
		return assets, err
	}
	for i := range outputs {
		assets[outputs[i].AssetId] = outputs[i].Amount
	}
	return assets, nil
}
