package validation

import (
	"github.com/kianooshaz/clean_service/contract"
	"github.com/kianooshaz/clean_service/param"
	"github.com/kianooshaz/clean_service/pkg/errors"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (v validate) LoginValidation(req *param.LoginRequestUser) contract.IServiceError {
	const section = errors.Section("validation.LoginValidation")

	if req.Email == "" {
		return errors.NewBadRequestError(section, "email is empty")
	}
	if !emailRegex.MatchString(req.Email) {
		return errors.NewBadRequestError(section, "invalid email")
	}
	if req.Password == "" {
		return errors.NewBadRequestError(section, "password is empty")
	}
	if len(req.Password) < 8 {
		return errors.NewBadRequestError(section, "password is less than 8 characters")
	}
	return nil
}

func (v validate) EmailAndPasswordValidation(user *param.EntryUser) contract.IServiceError {
	const section = errors.Section("validation.EmailAndPasswordValidation")

	if user.Email == "" {
		return errors.NewBadRequestError(section, "email is empty")
	}
	if !emailRegex.MatchString(user.Email) {
		return errors.NewBadRequestError(section, "invalid email")
	}
	if user.Password == "" {
		return errors.NewBadRequestError(section, "password is empty")
	}
	if len(user.Password) < 8 {
		return errors.NewBadRequestError(section, "password is less than 8 characters")
	}
	return nil
}
