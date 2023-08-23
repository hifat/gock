package main

import (
	"github.com/hifat/gock/internal/config"
	"github.com/hifat/gock/internal/database"
	"github.com/hifat/gock/internal/model"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadAppConfig()
	db, _ := database.NewPostgresConnection(*cfg)
	GormMigrate(db)
}

func GormMigrate(db *gorm.DB) {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	db.AutoMigrate(
		&model.Task{},
	)
}
