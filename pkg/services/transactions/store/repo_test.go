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
		}
		err := repo.Create(context.Background(), transaction)
		assert.NoError(t, err)
	})
}
