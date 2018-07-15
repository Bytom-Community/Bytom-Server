package db

type (
	Transactions struct {
		TxId           string `xorm:"tx_id"`
		BlockHash      string `xorm:"block_hash"`
		Amount         int64  `xorm:"amount"`
		Version        int64  `xorm:"version"`
		SerializedSize int64  `xorm:"serialized_size"`
		TimeRange      int64  `xorm:"timerange"`
	}
)

func (tx *Transactions) TableName() string {
	return "transactions"
}
