package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/middleware"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/labstack/echo/v4"
)

type PhotoControllerImpl struct {
	Usecase usecase.PhotoUsecase
	Route   *echo.Echo
}

func NewPhotoController(usecase usecase.PhotoUsecase, route *echo.Echo) PhotoController {
	controller := &PhotoControllerImpl{
		Usecase: usecase,
		Route:   route,
	}

	controller.route(route)
	return controller
}
func (photoControllerImpl *PhotoControllerImpl) route(echo *echo.Echo) {
	photosGroup := echo.Group("/photos")
	photosGroup.Use(middleware.Auth)
	photosGroup.POST("", photoControllerImpl.PostPhoto)
	photosGroup.GET("", photoControllerImpl.GetPhoto)
	photosGroup.PUT("/:photoId", photoControllerImpl.UpdatePhoto)
	photosGroup.DELETE("/:photoId", photoControllerImpl.DeletePhoto)
}

func (controller *PhotoControllerImpl) PostPhoto(ctx echo.Context) error {
	request := request.Photo{}

	err := json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.UserId, err = helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	photoResponse, err := controller.Usecase.PostPhoto(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusCreated,
		Message: "Photo have been successfully posted",
		Data:    photoResponse,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PhotoControllerImpl) GetPhoto(ctx echo.Context) error {
	photoResponse, err := controller.Usecase.GetPhoto(ctx.Request().Context())
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success get all photos",
		Data:    photoResponse,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PhotoControllerImpl) UpdatePhoto(ctx echo.Context) error {
	request := request.Photo{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.UserId, err = helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	request.Id, err = strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		return err
	}

	photoResponse, err := controller.Usecase.UpdatePhoto(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success update photo",
		Data:    photoResponse,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PhotoControllerImpl) DeletePhoto(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		return err
	}

	err = controller.Usecase.DeletePhoto(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Your photo has been successfully deleted",
	}

	return ctx.JSON(http.StatusOK, webResponse)
}
