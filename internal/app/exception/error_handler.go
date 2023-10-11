package exception

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/dihanto/gosnap/model/web/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func ErrorHandler(err error, ctx echo.Context) {
	if validationError(err, ctx) {
		return
	}

	if err == bcrypt.ErrMismatchedHashAndPassword {
		errPasswordDoNotMatch(err, ctx)
		return
	}
	if strings.Contains(err.Error(), "token is expired") {
		unauthorizedError(err, ctx)
		return
	}
	if err == sql.ErrNoRows {
		notFoundError(err, ctx)
		return
	}

	internalServerError(err, ctx)
}

func notFoundError(err error, ctx echo.Context) {
	webResponse := response.WebResponse{
		Status:  http.StatusNotFound,
		Message: "Not Found",
	}

	ctx.JSON(http.StatusNotFound, webResponse)
}
func unauthorizedError(err error, ctx echo.Context) {
	webResponse := response.WebResponse{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
	}

	ctx.JSON(http.StatusUnauthorized, webResponse)
}

func internalServerError(err interface{}, ctx echo.Context) {
	var dataError string
	if err != nil {
		dataError = err.(error).Error()
	}
	webResponse := response.WebResponse{
		Status:  http.StatusInternalServerError,
		Message: dataError,
	}

	ctx.JSON(http.StatusInternalServerError, webResponse)
}

func validationError(errs interface{}, ctx echo.Context) bool {
	exception, ok := errs.(validator.ValidationErrors)
	if ok {
		for _, err := range exception {
			var message string
			fieldName := strings.ToLower(err.Field())
			tag := err.ActualTag()
			switch tag {
			case "required":
				message = fmt.Sprintf("%s is required", fieldName)
			case "email":
				message = fmt.Sprintf("%s is not a valid email", fieldName)
			case "email_uniq":
				message = fmt.Sprintf("%s must be unique", fieldName)
			case "username_uniq":
				message = fmt.Sprintf("%s must be unique", fieldName)
			case "follow":
				message = fmt.Sprintf("%s you have already following", fieldName)
			case "min":
				message = fmt.Sprintf("%s must be at least %s characters long", fieldName, err.Param())
			case "gt":
				message = fmt.Sprintf("%s must be greater than %s", fieldName, err.Param())
			default:
				message = fmt.Sprintf("validation error for %s: %s", fieldName, tag)
			}
			webResponse := response.WebResponse{
				Status:  http.StatusBadRequest,
				Message: message,
			}
			ctx.JSON(http.StatusBadRequest, webResponse)
		}
		return true
	}
	return false
}

func errPasswordDoNotMatch(err error, ctx echo.Context) {
	webResponse := response.WebResponse{
		Status:  http.StatusBadRequest,
		Message: "Password does not match",
	}

	ctx.JSON(http.StatusBadRequest, webResponse)
}
