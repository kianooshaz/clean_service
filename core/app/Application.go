package app

import (
	"github.com/kianooshaz/clean_service/core/config"
	"github.com/kianooshaz/clean_service/core/pkg/logs"
	"github.com/kianooshaz/clean_service/core/repository"
	"github.com/kianooshaz/clean_service/core/server"
	"github.com/kianooshaz/clean_service/core/service"
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
	userServ := service.NewUserService(repository.NewUserRepository())
	httpServer := server.NewHttpServer(oCfg, userServ)
	e.Logger.Fatal(httpServer.Start(oCfg.Server.Port))
}
