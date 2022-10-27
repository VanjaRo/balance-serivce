package servmocs

import "context"

type MockTransactionsService struct {
	DepositErr error
	FreezeErr  error
	ApplyErr   error
}

func (s *MockTransactionsService) Deposit(ctx context.Context, userId string, amount float64) error {
	return s.DepositErr
}
func (s *MockTransactionsService) Freeze(ctx context.Context, userId, orderId, serviceId string, amount float64) error {
	return s.FreezeErr
}
func (s *MockTransactionsService) Apply(ctx context.Context, serviceId, orderId string) error {
	return s.ApplyErr
}
