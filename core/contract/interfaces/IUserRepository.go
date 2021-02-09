package interfaces

import "github.com/kianooshaz/clean_service/core/entity"

type IUserRepository interface {
	Create(user *entity.User) (*entity.User, IServiceError)
	Get(id int) (*entity.User, IServiceError)
	Update(user *entity.User) (*entity.User, IServiceError)
	Delete(id int) IServiceError
	FindAll() ([]entity.User, IServiceError)
}
