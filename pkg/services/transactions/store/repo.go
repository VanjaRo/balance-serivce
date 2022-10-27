package store

import (
	"context"

	"github.com/VanjaRo/balance-serivce/pkg/services/transactions"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type transactionRepo struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) transactions.Repo {
	return &transactionRepo{
		DB: db,
	}
}
func (tr *transactionRepo) Create(ctx context.Context, transaction transactions.Transaction) error {
	transaction.Id = uuid.New().String()
	return tr.DB.Create(&transaction).Error
}

func (tr *transactionRepo) Migrate() error {
	return tr.DB.AutoMigrate(&transactions.Transaction{})
}
