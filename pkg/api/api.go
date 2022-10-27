package api

import (
	"context"

	"github.com/VanjaRo/balance-serivce/pkg/db"
	transactionsTransport "github.com/VanjaRo/balance-serivce/pkg/services/transactions/transport"
	usersTransport "github.com/VanjaRo/balance-serivce/pkg/services/users/transport"
	"github.com/VanjaRo/balance-serivce/pkg/utils/log"
	"github.com/gin-gonic/gin"
)

type Config struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	RunMigration bool

	AppHost string
	AppPort string
}

func Start(cfg *Config) {
	ctx := context.Background()
	conn, err := db.InitDB(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	if err != nil {
		log.Error(ctx, "unable to establish a database connection: %s", err.Error())
	}

	router := gin.New()

	usersTransport.ActivateHandlers(router, conn)
	transactionsTransport.ActivateHandlers(router, conn)

	if err := router.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
		log.Error(ctx, "unable to start the server: %s", err.Error())
	}
}
