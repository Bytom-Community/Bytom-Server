package db

type (
	TxOutputs struct {
		//Id              int    `xorm:"<-"`
		TxId            string `xorm:"tx_id" json:"-"`
		BlockHash       string `xorm:"block_hash" json:"-"`
		Address         string `xorm:"address" json:"address"`
		AssetId         string `xorm:"asset_id" json:"asset_id"`
		Amount          uint64 `xorm:"amount" json:"amount"`
		AssetDefinition string `xorm:"asset_definition" json:"-"`
		OutputId        string `xorm:"output_id" json:"output_id"`
		Type            string `xorm:"type" json:"type"`
	}
)

func (t *TxOutputs) TableName() string {
	return "tx_outputs"
}
