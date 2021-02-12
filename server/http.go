package server

import (
	"fmt"
	"github.com/kianooshaz/clean_service/config"
	"github.com/kianooshaz/clean_service/contract"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

type httpServer struct {
	handlers *handlers
	user     *echo.Group
}

func NewHttpServer(cfg config.Config, user contract.IUserService) *httpServer {
	userRoute := e.Group("/users")
	return &httpServer{
		handlers: NewHandlers(cfg, user),
		user:     userRoute,
	}
}

func (h *httpServer) Start(port int) error {
	h.user.POST("", h.handlers.Create)
	h.user.GET("/:id", h.handlers.Get)
	h.user.PUT("", h.handlers.Update)
	h.user.PATCH("", h.handlers.Update)
	h.user.DELETE("/:id", h.handlers.Delete)
	h.user.GET("", h.handlers.FindAll)

	if port == 0 {
		port = 8081
	}

	return e.Start(fmt.Sprintf(":%d", port))
}