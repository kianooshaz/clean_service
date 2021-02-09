package interfaces

import (
	"github.com/kianooshaz/clean_service/core/contract/param"
)

type IUserService interface {
	Create(entry *param.EntryUser) (*param.PublicUser, IServiceError)
	Get(id int) (*param.PublicUser, IServiceError)
	Update(entry *param.EntryUser, isPartial bool) (*param.PublicUser, IServiceError)
	Delete(id int) IServiceError
	FindAll() ([]param.PublicUser, IServiceError)
}
