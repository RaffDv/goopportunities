package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	logger := GetLogger("[CONFIG]")
	var err error

	db, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("error initializing postgres: %v", err)
	}
	logger.Info("Config init")

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
