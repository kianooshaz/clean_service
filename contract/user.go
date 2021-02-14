package contract

import (
	"github.com/kianooshaz/clean_service/param"
)

type IUserService interface {
	Login(req *param.LoginRequestUser) (*param.UserTokens, IServiceError)

	Create(entry *param.EntryUser) (*param.PublicUser, IServiceError)
	Get(id int) (*param.PublicUser, IServiceError)
	Update(entry *param.EntryUser, isPartial bool) (*param.PublicUser, IServiceError)
	Delete(id int) IServiceError
	FindAll() ([]param.PublicUser, IServiceError)
}
