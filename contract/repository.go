package contract

import "github.com/kianooshaz/clean_service/entity"

type IUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, IServiceError)
	GetUserByID(id int) (*entity.User, IServiceError)
	UpdateUser(user *entity.User) (*entity.User, IServiceError)
	DeleteUserByID(id int) IServiceError
	FindAllUser() ([]entity.User, IServiceError)
}
