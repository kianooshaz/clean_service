package interfaces

import (
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	FindAll(c echo.Context) error
}
