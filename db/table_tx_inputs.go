package db

type (
	TxInputs struct {
		//Id              uint    `xorm:"<-"`
		TxId            string `xorm:"tx_id" json:"-"`
		BlockHash       string `xorm:"block_hash" json:"-"`
		Address         string `xorm:"address" json:"address"`
		AssetId         string `xorm:"asset_id" json:"asset_id"`
		Amount          uint64 `xorm:"amount" json:"amount"`
		AssetDefinition string `xorm:"asset_definition" json:"-"`
		SpentOutputId   string `xorm:"spent_output_id" json:"spent_output_id"`
		Type            string `xorm:"type" json:"type"`
	}
)

func (t *TxInputs) TableName() string {
	return "tx_inputs"
}
