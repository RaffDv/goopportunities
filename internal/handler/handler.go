package handler

import (
	"github.com/RaffDv/goopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func init() {
	var err error

	logger = config.GetLogger("[HANDLER]")
	db, err = config.InitializePostgres()

	if err != nil {
		logger.Errorf("error initializing postgres: %v", err.Error())
		return
	}

	logger.Info("Handler init")

}
