package exception

import (
	"net/http"

	"github.com/dihanto/gosnap/model/web"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	if validationError(err, c) {
		return
	}

	internalServerError(err, c)
}

func validationError(err interface{}, c echo.Context) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		webResponse := web.WebResponse{
			Status: http.StatusBadRequest,
			Data:   exception.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}
	return false
}

func internalServerError(err interface{}, c echo.Context) {
	webResponse := web.WebResponse{
		Status: http.StatusInternalServerError,
		Data:   err,
	}

	c.JSON(http.StatusInternalServerError, webResponse)
}
