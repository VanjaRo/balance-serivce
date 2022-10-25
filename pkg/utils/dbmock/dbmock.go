package dbmock

import (
	"fmt"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// inits docker container with DB
func InitMockDB() (*gorm.DB, func()) {
	// create a new pool for docker container
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}

	const (
		dbName = "test"
		passwd = "test"
	)

	runDockerOpt := &dockertest.RunOptions{
		Repository: "postgres", // image
		Tag:        "14",       // version
		Env:        []string{"POSTGRES_PASSWORD=" + passwd, "POSTGRES_DB=" + dbName},
	}

	fnConfig := func(config *docker.HostConfig) {
		config.AutoRemove = true                     // set AutoRemove to true so that stopped container goes away by itself
		config.RestartPolicy = docker.NeverRestart() // don't restart container
	}

	resource, err := pool.RunWithOptions(runDockerOpt, fnConfig)
	if err != nil {
		panic(err)
	}
	fnCleanup := func() {
		err := resource.Close()
		if err != nil {
			panic(err)
		}
	}
	conStr := fmt.Sprintf("host=localhost port=%s user=postgres dbname=%s password=%s sslmode=disable",
		resource.GetPort("5432/tcp"), // get port of localhost
		dbName,
		passwd,
	)
	var gdb *gorm.DB
	// retry until db server is ready
	err = pool.Retry(func() error {
		gdb, err = gorm.Open(postgres.Open(conStr))
		if err != nil {
			return err
		}
		db, err := gdb.DB()
		if err != nil {
			return err
		}
		return db.Ping()
	})
	if err != nil {
		panic(err)
	}
	return gdb, fnCleanup
}

func FillDBWithData(db *gorm.DB, data ...interface{}) {
	for _, d := range data {
		err := db.Create(d).Error
		if err != nil {
			panic(err)
		}
	}
}
