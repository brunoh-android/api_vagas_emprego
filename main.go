package main

import (
	"github.com/brunoh-android/api_vagas_emprego.git/config"
	"github.com/brunoh-android/api_vagas_emprego.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}
	router.Initialize()
}
