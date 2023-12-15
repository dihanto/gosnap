package controller

import "github.com/labstack/echo/v4"

type PhotoController interface {
	PostPhoto(ctx echo.Context) error
	GetPhoto(ctx echo.Context) error
	UpdatePhoto(ctx echo.Context) error
	DeletePhoto(ctx echo.Context) error
	GetPhotoById(ctx echo.Context) error
}
