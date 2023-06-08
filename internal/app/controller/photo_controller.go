package controller

import "github.com/labstack/echo/v4"

type PhotoController interface {
	PostPhoto(c echo.Context) error
	GetPhoto(c echo.Context) error
	UpdatePhoto(c echo.Context) error
	DeletePhoto(c echo.Context) error
}
