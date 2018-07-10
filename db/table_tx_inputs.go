package db

type (
	TxInputs struct {
		//Id              uint    `xorm:"<-"`
		TxId            string `xorm:"tx_id"`
		BlockHash       string `xorm:"block_hash"`
		Address         string `xorm:"address"`
		AssetId         string `xorm:"asset_id"`
		Amount          uint64 `xorm:"amount"`
		AssetDefinition string `xorm:"asset_definition"`
		SpentOutputId   string `xorm:"spent_output_id"`
		Type            string `xorm:"type"`
	}
)

func (t *TxInputs) TableName() string {
	return "tx_inputs"
}
