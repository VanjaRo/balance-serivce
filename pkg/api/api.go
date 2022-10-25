package api

import (
	"github.com/VanjaRo/balance-serivce/pkg/db"
	usersTransport "github.com/VanjaRo/balance-serivce/pkg/services/users/transport"
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
	// ctx := context.Background()
	conn, err := db.InitDB(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	if err != nil {
		// log.Error(ctx, "unable to establish a database connection: %s", err.Error())
	}

	router := gin.New()

	usersTransport.ActivateHandlers(router, conn)
	if err := router.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
		// log.Error(ctx, "unable to start the server: %s", err.Error())
	}
}
