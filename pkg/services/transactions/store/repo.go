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

func (tr *transactionRepo) GetTrByOrderAndServiceIds(ctx context.Context, orderId, serviceId string) (transactions.Transaction, error) {
	var transaction transactions.Transaction
	err := tr.DB.Where("order_id = ? AND service_id = ?", orderId, serviceId).First(&transaction).Error
	if err != nil {
		return transactions.Transaction{}, transactions.ErrTransactionNotFound
	}
	return transaction, nil
}

func (tr *transactionRepo) Migrate() error {
	return tr.DB.AutoMigrate(&transactions.Transaction{})
}
