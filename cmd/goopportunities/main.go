package main

import (
	"fmt"

	"github.com/RaffDv/goopportunities/config"
	"github.com/RaffDv/goopportunities/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	fmt.Println("INIT")

	logger = config.GetLogger("[ MAIN ]")
	logger.Info("init main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	router.Initialize()

}
