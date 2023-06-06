package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	UserRegister(c echo.Context) error
	UserLogin(c echo.Context) error
	UserUpdate(c echo.Context) error
	UserDelete(c echo.Context) error
}
