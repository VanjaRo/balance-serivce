package store

import (
	"context"
	"testing"

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
	db.AutoMigrate(&users.User{})
	// create user
	t.Run("create user", func(t *testing.T) {
		repo := NewUserRepo(db)
		user := users.User{
			Id:      "1",
			Balance: 100,
		}
		id, err := repo.Create(context.Background(), user)
		assert.NoError(t, err)
		assert.NotEmpty(t, id)
	})
	// get user
	t.Run("get user", func(t *testing.T) {
		repo := NewUserRepo(db)
		user, err := repo.Get(context.Background(), "1")
		assert.NoError(t, err)
		assert.NotEmpty(t, user)
	})
	// get user not found
	t.Run("get user not found", func(t *testing.T) {
		repo := NewUserRepo(db)
		_, err := repo.Get(context.Background(), "2")
		assert.Error(t, err)
		assert.Equal(t, users.ErrUserNotFound, err)
	})
	// get all users
	t.Run("get all users", func(t *testing.T) {
		repo := NewUserRepo(db)
		users, err := repo.GetAll(context.Background(), 10, 0)
		assert.NoError(t, err)
		assert.NotEmpty(t, users)
	})
	// get user balance
	t.Run("get user balance", func(t *testing.T) {
		repo := NewUserRepo(db)
		balance, err := repo.GetUserBalance(context.Background(), "1")
		assert.NoError(t, err)
		assert.NotEmpty(t, balance)
	})
	// get user balance not found
	t.Run("get user balance not found", func(t *testing.T) {
		repo := NewUserRepo(db)
		_, err := repo.GetUserBalance(context.Background(), "3")
		assert.Error(t, err)
		assert.Equal(t, users.ErrUserNotFound, err)
	})
}
