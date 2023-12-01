package main

import (
	"github.com/RaffDv/goopportunities/config"
	"github.com/RaffDv/goopportunities/internal/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("[ MAIN ]")

	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	router.Initialize()

}
