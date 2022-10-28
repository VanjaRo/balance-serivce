package transactions

import "errors"

var (
	ErrNegativeOrZeroAmount = errors.New("amount must be positive")

	ErrTransactionNotFound = errors.New("transaction not found")

	ErrCantFreezeSameTransactionTwice = errors.New("can't freeze the same transaction twice")

	ErrAmountsDontMatch = errors.New("amounts of the frozen and input transactions don't match")

	ErrEmptyId = errors.New("user_id, service_id and order_id can't be empty")

	ErrCantApplyNotFrozenTransaction = errors.New("can't apply not frozen transaction")

	ErrTransactionAlreadyApplied = errors.New("transaction already applied")

	ErrCantRevertNotFrozenTransaction = errors.New("can't revert not frozen transaction")

	ErrTransactionQuery = errors.New("requested transactions could not be retrieved base on the given criteria")
)
