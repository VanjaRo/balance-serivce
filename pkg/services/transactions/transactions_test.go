package transactions

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type repoMock struct {
	CreateError error

	GetResult Transaction
	GetError  error

	UpdateErr error

	DeleteErr error

	GetApplTrsResult []Transaction
	GetApplTrsError  error

	ExportResult []ServicesStat
	ExportError  error
}

func (r *repoMock) Create(ctx context.Context, transaction Transaction) error {
	return r.CreateError
}

func (r *repoMock) GetTrByOrderAndServiceIds(ctx context.Context, orderId, serviceId string) (Transaction, error) {
	return r.GetResult, r.GetError
}

func (r *repoMock) UpdateTrStatus(ctx context.Context, t Transaction) error {
	return r.UpdateErr
}

func (r *repoMock) DeleteTr(ctx context.Context, t Transaction) error {
	return r.DeleteErr
}

func (r *repoMock) GetTrsByUserId(ctx context.Context, userId string, limit, offset int, sortConf *SortConfig) ([]Transaction, error) {
	return r.GetApplTrsResult, r.GetApplTrsError
}

func (r *repoMock) GetServicesStatsWithinYearMonth(ctx context.Context, year, month int) ([]ServicesStat, error) {
	return r.ExportResult, r.ExportError
}
func TestServiceDeposit(t *testing.T) {
	id := uuid.New().String()
	tests := map[string]struct {
		repo Repo
		err  error
	}{
		"Happy path": {
			repo: &repoMock{
				CreateError: nil,
			},
			err: nil,
		},
		"DB error": {
			repo: &repoMock{
				CreateError: fmt.Errorf("DB error"),
			},
			err: fmt.Errorf("DB error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			service := NewTransactionService(test.repo)
			err := service.Deposit(context.Background(), id, 100)

			assert.Equal(t, test.err, err)
		})
	}
}
