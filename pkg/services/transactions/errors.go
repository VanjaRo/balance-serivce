package transactions

import "errors"

var (
	ErrNegativeOrZeroAmount = errors.New("amount must be positive")

	ErrTransactionNotFound = errors.New("transaction not found")

	ErrCantFreezeSameTransactionTwice = errors.New("can't freeze the same transaction twice")

	ErrAmountsDontMatch = errors.New("amounts of frozen and aply don't match")

	ErrEmptyId = errors.New("user_id, service_id and order_id can't be empty")

	ErrCantApplyNotFrozenTransaction = errors.New("can't apply not frozen transaction")

	ErrTransactionAlreadyApplied = errors.New("transaction already applied")
)
