package store

import (
	"context"
	"fmt"

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

func (tr *transactionRepo) GetTrsByUserId(ctx context.Context, userId string, limit, offset int, sortConf *transactions.SortConfig) ([]transactions.Transaction, error) {
	var ts []transactions.Transaction
	query := tr.DB.Where("user_id = ?", userId).Limit(limit).Offset(offset)
	if sortConf != nil {
		if sortConf.ByAmountAsc {
			query = query.Order("amount asc")
		}
		if sortConf.ByAmountDesc {
			query = query.Order("amount desc")
		}
		if sortConf.ByDateAsc {
			query = query.Order("updated_at asc")
		}
		if sortConf.ByDateDesc {
			query = query.Order("updated_at desc")
		}
	}
	err := query.Find(&ts).Error
	if err != nil {
		log.Error(ctx, "error while getting applied transactions by user id: %s", userId)
		return []transactions.Transaction{}, transactions.ErrTransactionQuery
	}
	return ts, nil

}

func (tr *transactionRepo) GetServicesStatsWithinYearMonth(ctx context.Context, year, month int) ([]transactions.ServicesStat, error) {
	// should return list of mappings service_id -> amount
	servicesStats := []transactions.ServicesStat{}
	// select tranactions within given month and year
	SQLQeryForPeriod := fmt.Sprintf(
		"SELECT service_id, sum(amount) FROM transactions WHERE EXTRACT(YEAR FROM updated_at) = %d AND EXTRACT(MONTH FROM updated_at) = %d AND state = '%s' AND type = '%s' GROUP BY service_id",
		year,
		month,
		transactions.TRANSACTION_STATE_APPLIED,
		transactions.TRANSACTION_TYPE_WITHDRAWAL)
	err := tr.DB.Raw(SQLQeryForPeriod).Scan(&servicesStats).Error
	if err != nil {
		log.Error(ctx, "error while exporting transactions within month")
		return nil, err
	}
	log.Info(ctx, "exporting service stat", servicesStats)
	return servicesStats, nil
}

func (tr *transactionRepo) Migrate() error {
	return tr.DB.AutoMigrate(&transactions.Transaction{})
}
