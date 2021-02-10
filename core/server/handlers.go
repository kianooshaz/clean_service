package server

import (
	"github.com/kianooshaz/clean_service/core/config"
	"github.com/kianooshaz/clean_service/core/contract/interfaces"
)

type handlers struct {
	config config.Config
	user   interfaces.IUserService
}

func NewHandlers(config config.Config, user interfaces.IUserService) *handlers {
	return &handlers{config: config, user: user}
}
