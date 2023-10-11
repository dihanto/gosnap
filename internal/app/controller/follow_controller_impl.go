package controller

import (
	"net/http"
	"strings"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/middleware"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/labstack/echo/v4"
)

type FollowControllerImpl struct {
	Usecase usecase.FollowUsecase
	Route   *echo.Echo
}

func NewFollowControllerImpl(usecase usecase.FollowUsecase, route *echo.Echo) FollowController {
	controller := &FollowControllerImpl{
		Usecase: usecase,
		Route:   route,
	}

	controller.route(route)
	return controller
}

func (controller *FollowControllerImpl) route(echo *echo.Echo) {
	followGroup := echo.Group("/follows")
	followGroup.Use(middleware.Auth)
	followGroup.POST("/:username", controller.FollowUser)
	followGroup.DELETE("/:username", controller.UnfollowUser)
}

func (controller *FollowControllerImpl) FollowUser(ctx echo.Context) (err error) {
	request := request.Follow{}
	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.FollowerUsername, err = helper.GetUsernameFromToken(tokenString)
	if err != nil {
		return
	}

	request.TargetUsername = ctx.Param("username")

	follow, err := controller.Usecase.FollowUser(ctx.Request().Context(), request)
	if err != nil {
		return
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success follow " + request.TargetUsername,
		Data:    follow,
	}

	err = ctx.JSON(http.StatusOK, webResponse)
	if err != nil {
		return
	}

	return
}

func (controller *FollowControllerImpl) UnfollowUser(ctx echo.Context) (err error) {
	request := request.Follow{}
	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.FollowerUsername, err = helper.GetUsernameFromToken(tokenString)
	if err != nil {
		return
	}

	request.TargetUsername = ctx.Param("username")

	follow, err := controller.Usecase.UnFollowUser(ctx.Request().Context(), request)
	if err != nil {
		return
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Unfollow " + request.TargetUsername + " success",
		Data:    follow,
	}

	err = ctx.JSON(http.StatusOK, webResponse)
	if err != nil {
		return
	}
	return
}
