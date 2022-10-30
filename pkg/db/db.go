package db

import (
	"fmt"

	"github.com/VanjaRo/balance-serivce/pkg/services/transactions"
	"github.com/VanjaRo/balance-serivce/pkg/services/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(host, port, user, password, dbName string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	var db *gorm.DB
	var err error
	// retry until db server is ready
	for {
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			continue
		}
		break
	}

	if err != nil {
		return nil, err
	}
	// Migrate the schemas
	db.AutoMigrate(&users.User{}, &transactions.Transaction{})

	return db, nil
}
