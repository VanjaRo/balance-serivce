package transactions

import "context"

// Repo defines the DB level interaction of transactions
type Repo interface {
	Create(ctx context.Context, transaction Transaction) error
	GetTrByOrderAndServiceIds(ctx context.Context, orderId, serviceId string) (Transaction, error)
	UpdateTrStatus(ctx context.Context, t Transaction) error
	DeleteTr(ctx context.Context, t Transaction) error
	GetTrsByUserId(ctx context.Context, userId string, limit, offset int, sortConf *SortConfig) ([]Transaction, error)
}

// Service defines the business logic of users
type Service interface {
	Deposit(ctx context.Context, userId string, amount float64) error
	Freeze(ctx context.Context, userId, orderId, service_id string, amount float64) error
	Apply(ctx context.Context, userId, orderId, service_id string, amount float64) error
	Revert(ctx context.Context, userId, orderId, service_id string, amount float64) error
	GetUserStat(ctx context.Context, userId string, limit, offset int, sortConf *SortConfig) ([]Transaction, error)
}

type transaction struct {
	repo Repo
}

func NewTransactionService(repo Repo) Service {
	return &transaction{
		repo: repo,
	}
}
func (t *transaction) Deposit(ctx context.Context, userId string, amount float64) error {
	// check if the amount is positive
	if amount <= 0 {
		return ErrNegativeOrZeroAmount
	}
	// create a deposit transaction
	transaction := Transaction{
		UserId:    userId,
		Amount:    amount,
		State:     "APPLIED",
		IsDeposit: true,
	}
	err := t.repo.Create(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
func (t *transaction) Freeze(ctx context.Context, userId, orderId, serviceId string, amount float64) error {
	// check if the amount is positive
	if amount <= 0 {
		return ErrNegativeOrZeroAmount
	}
	// check that the ids are not empty
	if userId == "" || serviceId == "" || orderId == "" {
		return ErrEmptyId
	}
	// it looks like we can uniquely identify a transaction by the orderId and serviceId
	_, err := t.repo.GetTrByOrderAndServiceIds(ctx, orderId, serviceId)
	// if the transaction already exists, we can't freeze it again
	if err != nil {
		// create a freeze transaction
		transaction := Transaction{
			UserId:    userId,
			OrderId:   orderId,
			ServiceId: serviceId,
			Amount:    amount,
			State:     TRANSACTION_STATE_FROZEN,
			IsDeposit: false,
		}
		err := t.repo.Create(ctx, transaction)
		if err != nil {
			return err
		}
		return nil
	}

	return ErrCantFreezeSameTransactionTwice
}
func (t *transaction) Apply(ctx context.Context, userId, orderId, serviceId string, amount float64) error {
	// apply only frozen transactions
	// check that the ids are not empty
	if userId == "" || serviceId == "" || orderId == "" {
		return ErrEmptyId
	}
	// get frozen transaction by orderId and serviceId
	tr, err := t.repo.GetTrByOrderAndServiceIds(ctx, orderId, serviceId)
	if err != nil {
		return err
	}
	// check that the transaction is frozen
	if tr.State != TRANSACTION_STATE_FROZEN {
		return ErrCantApplyNotFrozenTransaction
	}
	// check that amount is the same
	if tr.Amount != amount {
		return ErrAmountsDontMatch
	}

	// update the transaction state to applied
	tr.State = TRANSACTION_STATE_APPLIED
	err = t.repo.UpdateTrStatus(ctx, tr)
	if err != nil {
		return err
	}
	return nil
}

func (t *transaction) Revert(ctx context.Context, userId, orderId, serviceId string, amount float64) error {
	// revert only frozen transactions
	// check that the ids are not empty
	if userId == "" || serviceId == "" || orderId == "" {
		return ErrEmptyId
	}
	// get frozen transaction by orderId and serviceId
	tr, err := t.repo.GetTrByOrderAndServiceIds(ctx, orderId, serviceId)
	if err != nil {
		return err
	}
	// check that the transaction is frozen
	if tr.State != TRANSACTION_STATE_FROZEN {
		return ErrCantRevertNotFrozenTransaction
	}
	// check that amount is the same
	if tr.Amount != amount {
		return ErrAmountsDontMatch
	}
	// try to delete the transaction
	err = t.repo.DeleteTr(ctx, tr)
	if err != nil {
		return err
	}
	return nil

}

func (t *transaction) GetUserStat(ctx context.Context, userId string, limit, offset int, sortConf *SortConfig) ([]Transaction, error) {
	return t.repo.GetTrsByUserId(ctx, userId, limit, offset, sortConf)
}

func (t *transaction) TimePeriodStatExport(ctx context.Context, userId string, limit, offset int, sortedByDate, sortedByTime bool) error {
	// cancel only frozen transactions
	// check that the ids are not empty
	return nil
}
