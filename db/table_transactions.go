package db

type (
	Transactions struct {
		TxId           string `xorm:"tx_id"`
		BlockHash      string `xorm:"block_hash"`
		Amount         uint64 `xorm:"amount"`
		SerializedSize uint64 `xorm:"serialized_size"`
		TimeRange      uint64 `xorm:"timerange"`
	}
)

func (tx *Transactions) TableName() string {
	return "transactions"
}
