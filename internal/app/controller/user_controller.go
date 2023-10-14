package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	UserRegister(ctx echo.Context) error
	UserLogin(ctx echo.Context) error
	UserUpdate(ctx echo.Context) error
	UserDelete(ctx echo.Context) error
	FindUser(ctx echo.Context) error
	FindAllUser(ctx echo.Context) error
}
