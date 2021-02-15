package psql

import (
	er "errors"
	"github.com/kianooshaz/clean_service/contract"
	"github.com/kianooshaz/clean_service/entity"
	"github.com/kianooshaz/clean_service/pkg/errors"
	"gorm.io/gorm"
	"log"
)

func (p psql) CreateUser(user *entity.User) (*entity.User, contract.IServiceError) {
	const section = errors.Section("psql.CreateUser")
	if err := p.db.Create(user).Error; err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError(section, "database error", err)
	}

	return user, nil
}

func (p psql) GetUserByID(id int) (*entity.User, contract.IServiceError) {
	const section = errors.Section("psql.GetUserByID")
	user := &entity.User{Base: entity.Base{ID: uint(id)}}

	if err := p.db.First(user).Error; err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NewNotFound(section, "user not found")
		}

		log.Println(err)
		return nil, errors.NewInternalServerError(section, "database error", err)
	}

	return user, nil
}

func (p psql) GetUserByEmail(email string) (*entity.User, contract.IServiceError) {
	const section = errors.Section("psql.GetUserByEmail")
	user := &entity.User{Email: email}

	if err := p.db.Where("email = ?", user.Email).First(user).Error; err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NewNotFound(section, "user not found")
		}

		log.Println(err)
		return nil, errors.NewInternalServerError(section, "database error", err)
	}

	return user, nil
}

func (p psql) UpdateUser(user *entity.User) (*entity.User, contract.IServiceError) {
	const section = errors.Section("psql.UpdateUser")

	if err := p.db.Save(user).Error; err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError(section, "database error", err)
	}

	return user, nil
}

func (p psql) DeleteUserByID(id int) contract.IServiceError {
	const section = errors.Section("psql.DeleteUserByID")

	user := &entity.User{Base: entity.Base{ID: uint(id)}}
	if err := p.db.Delete(user).Error; err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			return errors.NewNotFound(section, "user not found")
		}

		log.Println(err)
		return errors.NewInternalServerError(section, "database error", err)
	}
	return nil
}

func (p psql) FindAllUser() ([]entity.User, contract.IServiceError) {
	const section = errors.Section("psql.FindAllUser")

	var users []entity.User
	if err := p.db.Find(&users).Error; err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NewNotFound(section, "user not found")
		}

		log.Println(err)
		return nil, errors.NewInternalServerError(section, "database error", err)
	}
	return users, nil
}
