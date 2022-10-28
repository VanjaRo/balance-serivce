package servmocs

import (
	"context"

	"github.com/VanjaRo/balance-serivce/pkg/services/transactions"
)

type MockTransactionsService struct {
	DepositErr error

	FreezeErr error

	ApplyErr error

	RevertErr error

	GetUserTrsResult []transactions.Transaction
	GetUserTrsError  error
}

func (s *MockTransactionsService) Deposit(ctx context.Context, userId string, amount float64) error {
	return s.DepositErr
}
func (s *MockTransactionsService) Freeze(ctx context.Context, userId, orderId, serviceId string, amount float64) error {
	return s.FreezeErr
}
func (s *MockTransactionsService) Apply(ctx context.Context, userId, orderId, service_id string, amount float64) error {
	return s.ApplyErr
}

func (s *MockTransactionsService) Revert(ctx context.Context, userId, orderId, service_id string, amount float64) error {
	return s.ApplyErr
}
func (s *MockTransactionsService) GetUserTrs(ctx context.Context, userId string, limit, offset int, sortConf *transactions.SortConfig) ([]transactions.Transaction, error) {
	return s.GetUserTrsResult, s.GetUserTrsError
}
