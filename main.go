package main

import (
	"github.com/danielbrazrocha/goportunities/config"
	router "github.com/danielbrazrocha/goportunities/router"
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
