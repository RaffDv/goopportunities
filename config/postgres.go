package config

import (
	"github.com/RaffDv/goopportunities/internal/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("[postgres]")

	dsn := "host=172.24.0.2 user=postgres password=postgresql port=5432 dbname=goopportunities sslmode=disable "

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("connect error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schema.Opening{})
	if err != nil {
		logger.Errorf("AutoMigrate error: %v", err)
		return nil, err
	}

	return db, nil

}
