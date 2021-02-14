package user

import (
	"github.com/kianooshaz/clean_service/contract"
	"github.com/kianooshaz/clean_service/param"
)

func (s *userService) Login(user *param.LoginRequestUser) (*param.UserTokens, contract.IServiceError) {
	return nil, nil
}
