package validation

import (
	"github.com/kianooshaz/clean_service/contract"
)

type validate struct {
}

func NewValidate() contract.IUserValidation {
	return &validate{}
}
