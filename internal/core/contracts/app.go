package contracts

import (
	"core-api/internal/dbconnect"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type App struct {
	DB       *sqlx.DB    // provide db connection
	Router   *mux.Router // provide list routes for api
	Services *Services   // provide list service that called by handler
}

type Services struct {
	//LIST SERVICE THAT WILL BE USED FOR THIS API
	Soccer SoccerService
}

func (app *App) Initialize() {
	app.DB = dbconnect.CreateDBConnection()
	app.Router = mux.NewRouter()
}
