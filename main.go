package main

import (
	"github.com/becardine/gopportunities/config"
	"github.com/becardine/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	// init router
	router.Init()
}
