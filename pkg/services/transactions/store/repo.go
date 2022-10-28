package store

import (
	"context"

	"github.com/VanjaRo/balance-serivce/pkg/services/transactions"
	"github.com/VanjaRo/balance-serivce/pkg/utils/log"
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
	err := tr.DB.Where("service_id = ? AND order_id = ?", serviceId, orderId).First(&transaction).Error
	if err != nil {
		log.Error(ctx, "error while getting transaction by order id: %s and service id: %s", orderId, serviceId)
		return transactions.Transaction{}, transactions.ErrTransactionNotFound
	}
	return transaction, nil
}

func (tr *transactionRepo) UpdateTrStatus(ctx context.Context, t transactions.Transaction) error {

	result := tr.DB.Model(&transactions.Transaction{}).Where("user_id = ? AND service_id = ? AND order_id = ? AND amount = ?", t.UserId, t.ServiceId, t.OrderId, t.Amount).Updates(&t)
	if result.Error != nil {
		log.Error(ctx, "error while updating transaction status")
		return result.Error
	}
	// check if transaction was updated
	if result.RowsAffected == 0 {
		log.Error(ctx, "transaction was not updated")
		return transactions.ErrTransactionNotFound
	}
	return nil
}

func (tr *transactionRepo) DeleteTr(ctx context.Context, t transactions.Transaction) error {

	result := tr.DB.Model(&transactions.Transaction{}).Delete(&t)
	if result.Error != nil {
		log.Error(ctx, "error while deleting transaction status")
		return result.Error
	}
	// check if transaction was updated
	if result.RowsAffected == 0 {
		log.Error(ctx, "transaction was not deleted")
		return transactions.ErrTransactionNotFound
	}
	return nil
}

func (tr *transactionRepo) Migrate() error {
	return tr.DB.AutoMigrate(&transactions.Transaction{})
}
