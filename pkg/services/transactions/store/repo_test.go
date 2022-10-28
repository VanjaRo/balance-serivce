package store

import (
	"context"
	"testing"

	"github.com/VanjaRo/balance-serivce/pkg/services/transactions"
	"github.com/VanjaRo/balance-serivce/pkg/services/users"
	"github.com/VanjaRo/balance-serivce/pkg/utils/dbmock"
	"github.com/stretchr/testify/assert"
)

// fill it with test data
// test
// drop the container
func TestUserRepo(t *testing.T) {
	db, fnCleanup := dbmock.InitMockDB()
	defer fnCleanup()
	db.AutoMigrate(&users.User{}, &transactions.Transaction{})
	// create transaction
	t.Run("create transaction", func(t *testing.T) {
		repo := NewTransactionRepo(db)
		transaction := transactions.Transaction{
			OrderId:   "1",
			UserId:    "1",
			ServiceId: "1",
			Amount:    100,
			State:     transactions.TRANSACTION_STATE_FROZEN,
			IsDeposit: false,
		}
		err := repo.Create(context.Background(), transaction)
		assert.NoError(t, err)
	})
	// get transaction
	t.Run("get transaction", func(t *testing.T) {
		repo := NewTransactionRepo(db)
		transaction, err := repo.GetTrByOrderAndServiceIds(context.Background(), "1", "1")
		assert.NoError(t, err)
		assert.NotEmpty(t, transaction)
	})
	// update transaction
	t.Run("update transaction", func(t *testing.T) {
		repo := NewTransactionRepo(db)
		transaction, err := repo.GetTrByOrderAndServiceIds(context.Background(), "1", "1")
		assert.NoError(t, err)
		assert.NotEmpty(t, transaction)
		transaction.State = transactions.TRANSACTION_STATE_APPLIED
		err = repo.UpdateTrStatus(context.Background(), transaction)
		assert.NoError(t, err)
	})
	// export transactions within year-month
	t.Run("export transactions within year-month", func(t *testing.T) {
		repo := NewTransactionRepo(db)
		ts, err := repo.GetServicesStatsWithinYearMonth(context.Background(), 2022, 10)
		assert.NoError(t, err)
		assert.NotEmpty(t, ts)
	})
}
