package controller

import "github.com/labstack/echo/v4"

type CommentController interface {
	PostComment(ctx echo.Context) error
	GetComment(ctx echo.Context) error
	UpdateComment(ctx echo.Context) error
	DeleteComment(ctx echo.Context) error
}
