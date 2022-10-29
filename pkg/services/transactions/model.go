package transactions

import "time"

const (
	TRANSACTION_STATE_FROZEN  = "FROZEN"
	TRANSACTION_STATE_APPLIED = "APPLIED"
)

// State field represents the state of the transaction
// Possible values are:
// 1. "FROZEN"
// 2. "APPLIED"
type Transaction struct {
	Id        string    `gorm:"primaryKey; autoIncrement" json:"id"`
	OrderId   string    `json:"order_id"`
	UserId    string    `json:"user_id"`
	ServiceId string    `json:"service_id"`
	Amount    int       `json:"amount"`
	State     string    `json:"state"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

// For the deposit transaction, the service_id and order_id fields will be empty

// State field is a part of the model because
// making it a separate entity would increse the complexity of the DB structure
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

// Sorting struct is used to sort the transactions
type SortConfig struct {
	ByDateAsc  bool
	ByDateDesc bool

	ByAmountAsc  bool
	ByAmountDesc bool
}

type ServicesStat struct {
	ServiceId string `json:"service_id"`
	Sum       int    `json:"sum"`
}
