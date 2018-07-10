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
