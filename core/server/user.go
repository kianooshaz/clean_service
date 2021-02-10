package server

import (
	"github.com/kianooshaz/clean_service/core/contract/param"
	"github.com/kianooshaz/clean_service/core/pkg/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *handlers) Create(c echo.Context) error {

	user := &param.EntryUser{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
	}

	result, serErr := h.user.Create(user)
	if serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *handlers) Get(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("id should be number"))
	}

	result, serErr := h.user.Get(ID)
	if serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *handlers) Update(c echo.Context) error {
	user := &param.EntryUser{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body"))
	}

	result, serErr := h.user.Update(user, isPartial(c.Request()))
	if serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *handlers) Delete(c echo.Context) error {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("id should be number"))
	}

	if serErr := h.user.Delete(ID); serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (h *handlers) FindAll(c echo.Context) error {

	result, serErr := h.user.FindAll()
	if serErr != nil {
		return c.JSON(serErr.GetStatus(), serErr)
	}

	return c.JSON(http.StatusOK, result)
}

func isPartial(r *http.Request) bool {

	return r.Method == http.MethodPatch
}
