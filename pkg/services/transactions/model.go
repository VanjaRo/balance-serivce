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
	Amount    float64   `json:"amount"`
	State     string    `json:"state"`
	IsDeposit bool      `json:"type"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"update_time"`
}

// For the deposit transaction, the service_id and order_id fields will be empty

// State field is a part of the model because
// making it a separate entity would increse the complexity of the DB structure
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}
