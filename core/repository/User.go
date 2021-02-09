package repository

import (
	er "errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kianooshaz/clean_service/core/contract/interfaces"
	"github.com/kianooshaz/clean_service/core/entity"
	"github.com/kianooshaz/clean_service/core/utils/errors"
	"github.com/kianooshaz/clean_service/core/utils/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

type userRepository struct {
}

func init() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.ErrorLogger.Fatalln(err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		logs.ErrorLogger.Fatalln(err)
	}
}
func NewUserRepository() interfaces.IUserRepository {

	return &userRepository{}
}

func (u userRepository) Create(user *entity.User) (*entity.User, interfaces.IServiceError) {

	if err := db.Create(user).Error; err != nil {
		logs.WarningLogger.Println(err)
		return nil, errors.NewInternalServerError("database error", err)
	}

	return user, nil
}

func (u userRepository) Get(id int) (*entity.User, interfaces.IServiceError) {

	user := &entity.User{Base: entity.Base{ID: uint(id)}}

	if err := db.First(user).Error; err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NewNotFound("user not found")
		}

		logs.WarningLogger.Println(err)
		return nil, errors.NewInternalServerError("database error", err)
	}

	return user, nil
}

func (u userRepository) Update(user *entity.User) (*entity.User, interfaces.IServiceError) {
	fmt.Println("test repository") // todo delete
	if err := db.Save(user).Error; err != nil {
		logs.WarningLogger.Println(err)
		return nil, errors.NewInternalServerError("database error", err)
	}

	return user, nil
}

func (u userRepository) Delete(id int) interfaces.IServiceError {

	user := &entity.User{Base: entity.Base{ID: uint(id)}}
	if err := db.Delete(user).Error; err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			return errors.NewNotFound("user not found")
		}

		logs.WarningLogger.Println(err)
		return errors.NewInternalServerError("database error", err)
	}
	return nil
}

func (u userRepository) FindAll() ([]entity.User, interfaces.IServiceError) {
	var users []entity.User
	if err := db.Find(&users).Error; err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NewNotFound("user not found")
		}

		logs.WarningLogger.Println(err)
		return nil, errors.NewInternalServerError("database error", err)
	}
	return users, nil
}
