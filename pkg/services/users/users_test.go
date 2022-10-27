package users

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type repoMock struct {
	GetResult User
	GetError  error

	GetBalanceResult float64
	GetBalanceError  error

	GetAllResult []User
	GetAllError  error

	CreateResult string
	CreateError  error

	UpdateErr error
}

func (r *repoMock) Get(ctx context.Context, id string) (User, error) {
	return r.GetResult, r.GetError
}

func (r *repoMock) GetUserBalance(ctx context.Context, id string) (float64, error) {
	return r.GetBalanceResult, r.GetBalanceError
}

func (r *repoMock) GetAll(ctx context.Context, limit, offset int) ([]User, error) {
	return r.GetAllResult, r.GetAllError
}

func (r *repoMock) Create(ctx context.Context, u User) (string, error) {
	return r.CreateResult, r.CreateError
}

func (r *repoMock) UpdateUserBalance(ctx context.Context, userId string, balanceDiff float64) error {
	return r.UpdateErr
}

func TestServiceGet(t *testing.T) {
	id := uuid.New().String()
	tests := map[string]struct {
		repo   Repo
		result User
		err    error
	}{
		"Happy path": {
			repo: &repoMock{
				GetResult: User{Id: id, Balance: 100, Version: 1},
				GetError:  nil,
			},
			result: User{Id: id, Balance: 100, Version: 1},
			err:    nil,
		},
		"Not found from repo": {
			repo: &repoMock{
				GetResult: User{},
				GetError:  ErrUserNotFound,
			},
			result: User{},
			err:    ErrUserNotFound,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			service := NewUserService(test.repo)
			response, err := service.Get(context.Background(), id)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.result.Id, response.Id)
		})
	}
}

func TestServiceGetBalance(t *testing.T) {
	id := uuid.New().String()
	balance := 100.0
	tests := map[string]struct {
		repo   Repo
		result float64
		err    error
	}{
		"Happy path": {
			repo: &repoMock{
				GetBalanceResult: balance,
				GetBalanceError:  nil,
			},
			result: balance,
			err:    nil,
		},
		"Not found from repo": {
			repo: &repoMock{
				GetBalanceResult: 0,
				GetBalanceError:  ErrUserNotFound,
			},
			result: 0,
			err:    ErrUserNotFound,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			service := NewUserService(test.repo)
			response, err := service.GetBalance(context.Background(), id)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.result, response)
		})
	}
}

func TestServiceCreate(t *testing.T) {
	id := uuid.New().String()
	balance := 100.0
	u := User{Id: id, Balance: balance}

	tests := map[string]struct {
		repo   Repo
		result string
		input  User
		err    error
	}{
		"Happy path": {
			repo: &repoMock{
				CreateResult: id,
				CreateError:  nil,
				GetResult:    User{},
				GetError:     ErrUserNotFound,
			},
			input:  u,
			result: id,
			err:    nil,
		},
		"User with that index already exists": {
			repo: &repoMock{
				GetResult:    User{Id: id, Balance: 200},
				GetError:     nil,
				CreateResult: "",
				CreateError:  ErrUserAlreadyExists,
			},
			input:  u,
			result: "",
			err:    ErrUserAlreadyExists,
		},
		"Creation with negative balance": {
			repo: &repoMock{
				CreateResult: id,
				CreateError:  ErrNegativeBalance,
				GetResult:    User{},
				GetError:     ErrUserNotFound,
			},
			input:  User{Id: id, Balance: -100},
			result: "",
			err:    ErrNegativeBalance,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			service := NewUserService(test.repo)
			response, err := service.Create(context.Background(), test.input.Id, test.input.Balance)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.result, response)
		})
	}
}
