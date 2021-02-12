package server

import (
	"github.com/kianooshaz/clean_service/config"
	"github.com/kianooshaz/clean_service/contract"
)

type handlers struct {
	config config.Config
	user   contract.IUserService
}

func NewHandlers(config config.Config, user contract.IUserService) *handlers {
	return &handlers{config: config, user: user}
}
