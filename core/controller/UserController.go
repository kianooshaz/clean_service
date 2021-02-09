package controller

import (
	"github.com/kianooshaz/clean_service/core/contract/interfaces"
	"github.com/kianooshaz/clean_service/core/contract/param"
	"github.com/kianooshaz/clean_service/core/utils/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type userController struct {
	Service interfaces.IUserService
}

func NewUserController(service interfaces.IUserService) interfaces.IUserController {
	return &userController{
		Service: service,
	}
}

func (u userController) Create(c echo.Context) error {

	user := &param.EntryUser{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
	}

	result, serErr := u.Service.Create(user)
	if serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusCreated, result)
}

func (u userController) Get(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("id should be number"))
	}

	result, serErr := u.Service.Get(ID)
	if serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusOK, result)
}

func (u userController) Update(c echo.Context) error {
	user := &param.EntryUser{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
	}

	result, serErr := u.Service.Update(user, isPartial(c.Request()))
	if serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusOK, result)
}

func (u userController) Delete(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("id should be number"))
	}

	if serErr := u.Service.Delete(ID); serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (u userController) FindAll(c echo.Context) error {

	result, serErr := u.Service.FindAll()
	if serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusOK, result)
}

func isPartial(r *http.Request) bool {

	return r.Method == http.MethodPatch
}
