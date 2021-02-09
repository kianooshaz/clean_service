package app

import (
	"github.com/joho/godotenv"
	"github.com/kianooshaz/clean_service/core/controllers"
	"github.com/kianooshaz/clean_service/core/repository"
	"github.com/kianooshaz/clean_service/core/services"
	"github.com/kianooshaz/clean_service/core/utils/logs"
	"github.com/labstack/echo/v4"
	"os"
)

var e = echo.New()

func StartApplication() {
	if err := godotenv.Load(); err != nil {
		logs.ErrorLogger.Fatalln("Error loading .env file")
	}
	userRouting(controllers.NewUserController(services.NewUserService(repository.NewUserRepository())))
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))
}
