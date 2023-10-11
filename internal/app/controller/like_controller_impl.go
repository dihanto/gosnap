package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/labstack/echo/v4"
)

type LikeControllerImpl struct {
	Usecase usecase.LikeUsecase
	Route   *echo.Echo
}

func NewLikeController(usecase usecase.LikeUsecase, route *echo.Echo) LikeController {
	controller := &LikeControllerImpl{
		Usecase: usecase,
		Route:   route,
	}

	controller.route(route)
	return controller
}

func (likeControllerImpl *LikeControllerImpl) route(echo *echo.Echo) {
	photosGroup := echo.Group("/photos")
	photosGroup.POST("/:photoId/likes", likeControllerImpl.LikePhoto)
	photosGroup.GET("/:photoId/likes", likeControllerImpl.IsLikePhoto)
	photosGroup.DELETE("/:photoId/unlikes", likeControllerImpl.UnlikePhoto)
}

func (controller *LikeControllerImpl) LikePhoto(ctx echo.Context) error {
	idString := ctx.Param("photoId")
	photoId, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	likeRequest := request.Like{
		PhotoId: photoId,
		UserId:  userId,
	}

	like, err := controller.Usecase.LikePhoto(ctx.Request().Context(), likeRequest)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Like photo success",
		Data:    like,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}

func (controller *LikeControllerImpl) UnlikePhoto(ctx echo.Context) error {
	idString := ctx.Param("photoId")
	photoId, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	likeRequest := request.Like{
		PhotoId: photoId,
		UserId:  userId,
	}

	like, err := controller.Usecase.UnlikePhoto(ctx.Request().Context(), likeRequest)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success unlike photo",
		Data:    like,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}

func (controller *LikeControllerImpl) IsLikePhoto(ctx echo.Context) error {
	idString := ctx.Param("photoId")
	photoId, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	likeRequest := request.Like{
		PhotoId: photoId,
		UserId:  userId,
	}

	like, err := controller.Usecase.IsLikePhoto(ctx.Request().Context(), likeRequest)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    like,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}
