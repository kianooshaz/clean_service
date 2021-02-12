package api

import (
	"github.com/kianooshaz/clean_service/config"
	"github.com/kianooshaz/clean_service/interactor"
	"github.com/kianooshaz/clean_service/pkg/logs"
	"github.com/kianooshaz/clean_service/repository"
	"github.com/kianooshaz/clean_service/server"
	"github.com/labstack/echo/v4"
	"os"
)

var e = echo.New()

var oCfg config.Config
var dbUri string

func init() {
	cfgPath := "./core/config/config.yml"

	if os.Getenv("CLEAN_SERVICE_CONFIGS_PATH") != "" {
		cfgPath = os.Getenv("CLEAN_SERVICE_CONFIGS_PATH")
	}

	err := config.ReadFile(&oCfg, cfgPath)
	if err != nil {
		logs.ErrorLogger.Fatalln(err)
	}

	err = config.ReadEnv(&oCfg)
	if err != nil {
		logs.ErrorLogger.Fatalln(err)
	}

	config.SetConfig(&oCfg)
}

func StartApplication() {
	userServ := interactor.NewUserService(repository.NewUserRepository())
	httpServer := server.NewHttpServer(oCfg, userServ)
	e.Logger.Fatal(httpServer.Start(oCfg.Server.Port))
}
