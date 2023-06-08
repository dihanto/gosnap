package controller

import "github.com/labstack/echo/v4"

type CommentController interface {
	PostComment(c echo.Context) error
	GetComment(c echo.Context) error
	UpdateComment(c echo.Context) error
	DeleteComment(c echo.Context) error
}
