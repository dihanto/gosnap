package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/dihanto/gosnap/model/web"
	"github.com/labstack/echo/v4"
)

type PhotoControllerImpl struct {
	Usecase usecase.PhotoUsecase
}

func NewPhotoController(usecase usecase.PhotoUsecase) PhotoController {
	return &PhotoControllerImpl{
		Usecase: usecase,
	}
}

func (controller *PhotoControllerImpl) PostPhoto(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	title := c.FormValue("title")
	caption := c.FormValue("caption")
	photoUrl := c.FormValue("photoUrl")

	request := web.Photo{
		Title:    title,
		Caption:  caption,
		PhotoUrl: photoUrl,
		UserId:   userId,
	}

	photoResponse, err := controller.Usecase.PostPhoto(c.Request().Context(), request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := web.WebResponse{
		Status: http.StatusCreated,
		Data:   photoResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *PhotoControllerImpl) GetPhoto(c echo.Context) error {
	photoResponse, err := controller.Usecase.GetPhoto(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   photoResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *PhotoControllerImpl) UpdatePhoto(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	id, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	title := c.FormValue("title")
	caption := c.FormValue("caption")
	photoUrl := c.FormValue("photoUrl")

	request := web.Photo{
		Title:    title,
		Caption:  caption,
		PhotoUrl: photoUrl,
		Id:       id,
		UserId:   userId,
	}

	photoResponse, err := controller.Usecase.UpdatePhoto(c.Request().Context(), request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   photoResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *PhotoControllerImpl) DeletePhoto(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = controller.Usecase.DeletePhoto(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   "Your photo has been successfully deleted",
	}

	return c.JSON(http.StatusOK, webResponse)
}
