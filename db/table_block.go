package db

type (
	Block struct {
		//Id                int    `xorm:"<-"`
		Hash              string `xorm:"hash"`
		Version           uint64 `xorm:"block_version"`
		Height            uint64 `xorm:"height"`
		PreviousBlockHash string `xorm:"previous_block_hash"`
		Timestamp         uint64 `xorm:"timestamp"`
		Nonce             uint64 `xorm:"nonce"`
		Bits              uint64 `xorm:"bits"`
		TxCount           int    `xorm:"tx_count"`
		TxIds             string `xorm:"tx_ids"`
	}
)

func (b *Block) TableName() string {
	return "block"
}
