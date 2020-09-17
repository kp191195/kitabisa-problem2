package main

import (
	"core-api/internal/core/contracts"
	"core-api/internal/core/handlers"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	app := &contracts.App{}
	app.Initialize()
	app.Services = initServices(app)
	handlers.Init(app.Services)
	initRoutes(app.Router)

	fmt.Println(fmt.Sprintf("> Starting your API Server on localhost:%s", viper.GetString("app.port")))
	http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("app.port")), app.Router)
}
