package controller

import "github.com/labstack/echo/v4"

type LikeController interface {
	LikePhoto(ctx echo.Context) error
	UnlikePhoto(ctx echo.Context) error
	IsLikePhoto(ctx echo.Context) error
}
