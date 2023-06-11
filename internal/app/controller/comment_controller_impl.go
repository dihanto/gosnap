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

type CommentControllerImpl struct {
	Usecase usecase.CommentUsecase
}

func NewCommentController(usecase usecase.CommentUsecase) CommentController {
	return &CommentControllerImpl{
		Usecase: usecase,
	}
}
func (controller *CommentControllerImpl) PostComment(c echo.Context) error {
	message := c.FormValue("message")
	photoIdString := c.FormValue("photoId")
	photoId, err := strconv.Atoi(photoIdString)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	request := web.PostComment{
		Message: message,
		PhotoId: photoId,
		UserId:  userId,
	}

	commentResponse, err := controller.Usecase.PostComment(c.Request().Context(), request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := web.WebResponse{
		Status: http.StatusCreated,
		Data:   commentResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *CommentControllerImpl) GetComment(c echo.Context) error {
	commentResponse, err := controller.Usecase.GetComment(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   commentResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *CommentControllerImpl) UpdateComment(c echo.Context) error {
	idString := c.Param("commentId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	message := c.FormValue("message")
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	request := web.UpdateComment{
		Id:      id,
		Message: message,
		UserId:  userId,
	}

	commentResponse, err := controller.Usecase.UpdateComment(c.Request().Context(), request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   commentResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *CommentControllerImpl) DeleteComment(c echo.Context) error {
	idString := c.Param("commentId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = controller.Usecase.DeleteComment(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   "Your comment has been successfully deleted",
	}

	return c.JSON(http.StatusOK, webResponse)
}
