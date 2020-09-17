package main

import (
	"core-api/internal/core/contracts"
	"core-api/internal/core/services/soccer"
)

func initServices(app *contracts.App) *contracts.Services {
	return &contracts.Services{
		// LIST SERVICE
		Soccer: soccer.Init(app),
	}
}
