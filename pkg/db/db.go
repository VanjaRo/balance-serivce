package db

import (
	"context"
	"fmt"

	"github.com/VanjaRo/balance-serivce/pkg/services/transactions"
	"github.com/VanjaRo/balance-serivce/pkg/services/users"
	"github.com/VanjaRo/balance-serivce/pkg/utils/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(ctx context.Context, host, port, user, password, dbName string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	var db *gorm.DB
	var err error
	// retry until db server is ready
	log.Info(ctx, "connecting to db...")
	for i := 0; i < 100000; i++ {

		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			continue
		}
		break
	}

	log.Info(ctx, "connected to db...")
	if err != nil {
		return nil, err
	}
	// Migrate the schemas
	db.AutoMigrate(&users.User{}, &transactions.Transaction{})

	return db, nil
}
