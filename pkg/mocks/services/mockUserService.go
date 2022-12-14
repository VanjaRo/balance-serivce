package servmocs

import (
	"context"

	"github.com/VanjaRo/balance-serivce/pkg/services/users"
)

type MockUsersService struct {
	GetResult users.User
	GetErr    error

	GetAllResult []users.User
	GetAllErr    error

	GetBalanceResult int
	GetBalanceError  error

	CreateResult string
	CreateErr    error

	UpdateErr error
}

func (s *MockUsersService) Get(ctx context.Context, id string) (users.User, error) {
	return s.GetResult, s.GetErr
}

func (s *MockUsersService) GetAll(ctx context.Context, limit, offset int) ([]users.User, error) {
	return s.GetAllResult, s.GetAllErr
}

func (s *MockUsersService) Create(ctx context.Context, userId string, balance int) (string, error) {
	return s.CreateResult, s.CreateErr
}

func (s *MockUsersService) GetBalance(ctx context.Context, id string) (int, error) {
	return s.GetBalanceResult, s.GetBalanceError
}

func (s *MockUsersService) UpdateUserBalance(ctx context.Context, userId string, balanceDiff int) error {
	return s.UpdateErr
}
