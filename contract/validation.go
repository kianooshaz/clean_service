package contract

import "github.com/kianooshaz/clean_service/param"

type IUserValidation interface {
	LoginValidation(req *param.LoginRequestUser) IServiceError
	EmailAndPasswordValidation(user *param.EntryUser) IServiceError
}
