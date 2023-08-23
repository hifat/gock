package database

import (
	"fmt"

	"github.com/hifat/gock/internal/config"

	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var NewPostgresDBSet = wire.NewSet(NewPostgresConnection)

func NewPostgresConnection(config config.AppConfig) (*gorm.DB, func()) {
	dbConifig := config.DB
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConifig.Host,
		dbConifig.Username,
		dbConifig.Password,
		dbConifig.Name,
		dbConifig.Port,
	)
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}

	cleanup := func() {
		sqlDB, err := conn.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	return conn, cleanup
}
