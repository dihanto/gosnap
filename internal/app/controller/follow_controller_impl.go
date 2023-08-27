package controller

import (
	"net/http"
	"strings"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/middleware"
	"github.com/dihanto/gosnap/internal/app/usecase"
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
}

func (controller *FollowControllerImpl) FollowUser(ctx echo.Context) (err error) {
	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	followerId, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return
	}

	username := ctx.Param("username")

	err = controller.Usecase.FollowUser(ctx.Request().Context(), followerId, username)
	if err != nil {
		return
	}

	webResponse := response.WebResponse{
		Status:  http.StatusCreated,
		Message: "Success follow " + username,
		Data:    nil,
	}

	err = ctx.JSON(http.StatusCreated, webResponse)
	if err != nil {
		return
	}

	return
}

func (controller *FollowControllerImpl) UnfollowUser(ctx echo.Context) (err error) {
	panic("not implemented") // TODO: Implement
}
