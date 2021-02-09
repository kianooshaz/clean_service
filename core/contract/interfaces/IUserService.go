package interfaces

import (
	"github.com/kianooshaz/clean_service/core/contract/params"
)

type IUserService interface {
	Create(entry *params.EntryUser) (*params.PublicUser, IServiceError)
	Get(id int) (*params.PublicUser, IServiceError)
	Update(entry *params.EntryUser, isPartial bool) (*params.PublicUser, IServiceError)
	Delete(id int) IServiceError
	FindAll() ([]params.PublicUser, IServiceError)
}
