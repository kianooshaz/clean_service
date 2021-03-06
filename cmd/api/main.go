package main

import (
	"fmt"
	"github.com/kianooshaz/clean_service/config"
	"github.com/kianooshaz/clean_service/entity"
	"github.com/kianooshaz/clean_service/interactor/auth"
	"github.com/kianooshaz/clean_service/interactor/user"
	"github.com/kianooshaz/clean_service/interactor/validation"
	"github.com/kianooshaz/clean_service/repository/psql"
	"github.com/kianooshaz/clean_service/server"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var e = echo.New()

var oCfg config.Config
var dbUri string

func init() {
	cfgPath := "./config/config.yml"

	if os.Getenv("CLEAN_SERVICE_CONFIGS_PATH") != "" {
		cfgPath = os.Getenv("CLEAN_SERVICE_CONFIGS_PATH")
	}

	err := config.ReadFile(&oCfg, cfgPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = config.ReadEnv(&oCfg)
	if err != nil {
		log.Fatalln(err)
	}

	config.SetConfig(&oCfg)
}

func init() {
	dbUri = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		oCfg.Database.Host,
		oCfg.Database.Username,
		oCfg.Database.Password,
		oCfg.Database.DBName,
		oCfg.Database.Port,
		oCfg.Database.SSLMode,
		oCfg.Database.Timezone,
	)
}

func main() {
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalln(err)
	}

	repo := psql.New(db)
	authService := auth.NewAuthService(oCfg)
	validateService := validation.NewValidate()
	userService := user.NewService(oCfg, repo, authService, validateService)
	httpService := server.NewHttpServer(oCfg, userService)
	e.Logger.Fatal(httpService.Start(oCfg.Server.Port))
}
