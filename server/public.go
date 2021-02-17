package server

import (
	"net/http"

	"github.com/kianooshaz/clean_service/param"
	"github.com/kianooshaz/clean_service/pkg/errors"
	"github.com/labstack/echo/v4"
)

func (h *handlers) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, &struct{ Message string }{
		Message: "Everything is good!",
	})
}

func (h *handlers) Login(c echo.Context) error {
	const section = errors.Section("server.Login")

	req := &param.LoginRequestUser{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError(section, "invalid json body"))
	}

	token, serErr := h.user.Login(req)
	if serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}
	return c.JSON(http.StatusCreated, token)
}
