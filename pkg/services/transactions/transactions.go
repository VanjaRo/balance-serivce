package transactions

import "context"

// Repo defines the DB level interaction of transactions
type Repo interface {
	Create(ctx context.Context, transaction Transaction) error
}

// Service defines the business logic of users
type Service interface {
	Deposit(ctx context.Context, userId string, amount float64) error
	Freeze(ctx context.Context, userId, orderId, service_id string, amount float64) error
	Apply(ctx context.Context, serviceId, orderId string) error
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
	// create a freeze transaction
	// transaction := Transaction{
	// 	UserId:    userId,
	// 	OrderId:   orderId,
	// 	ServiceId: serviceId,
	// 	Amount:    amount,
	// 	State:     "FROZEN",
	// 	IsDeposit: false,
	// }

	return nil
}
func (t *transaction) Apply(ctx context.Context, serviceId, orderId string) error {
	return nil
}
