package transactions

import "errors"

var (
	ErrNegativeOrZeroAmount = errors.New("amount must be positive")
)
