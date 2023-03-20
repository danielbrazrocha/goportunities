package main

import (
	"github.com/danielbrazrocha/gopportunities/config"
	router "github.com/danielbrazrocha/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	// Initialize Logger
	logger = config.GetLogger("main")

	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v ", err)
		return
	}

	//Initialize router
	router.Initialize()
}
