package controller

import "github.com/labstack/echo/v4"

type FollowController interface {
	FollowUser(ctx echo.Context) (err error)
	UnfollowUser(ctx echo.Context) (err error)
}
