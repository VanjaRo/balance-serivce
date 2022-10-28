package transactions

import "errors"

var (
	ErrNegativeOrZeroAmount = errors.New("amount must be positive")

	ErrTransactionNotFound = errors.New("transaction not found")

	ErrCantFreezeSameTransactionTwice = errors.New("can't freeze the same transaction twice")
)
