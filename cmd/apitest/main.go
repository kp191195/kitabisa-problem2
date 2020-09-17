package apitest

import (
	"core-api/internal/core/contracts"
	"core-api/internal/core/services/soccer"
	"core-api/internal/dbconnect"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
)

type TestApp struct {
	contracts.App
	MockController *gomock.Controller
}

func InitTestDep() string {
	// Boot core service
	str, _ := os.Getwd()
	return str
}

func InitializeTestApp(t *testing.T, confDir string) TestApp {
	ctrl := gomock.NewController(t)

	app := contracts.App{
		DB: dbconnect.CreateTestingDBConnection(confDir),
	}

	services := contracts.Services{
		Soccer: soccer.Init(&app),
	}

	app.Services = &services

	return TestApp{App: app, MockController: ctrl}

}
