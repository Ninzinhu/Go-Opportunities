package main

import (
	

	"github.com/Ninzinhu/Go-Opportunities/config"
	"github.com/Ninzinhu/Go-Opportunities/router"
)


var (
	logger config.Logger
)


func main(){

	logger= *config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}