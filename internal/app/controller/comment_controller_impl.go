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

type CommentControllerImpl struct {
	Usecase usecase.CommentUsecase
	Route   *echo.Echo
}

func NewCommentController(usecase usecase.CommentUsecase, route *echo.Echo) CommentController {
	controller := &CommentControllerImpl{
		Usecase: usecase,
		Route:   route,
	}
	controller.route(route)
	return controller
}
func (commentControllerImpl *CommentControllerImpl) route(e *echo.Echo) {
	commentsGroup := e.Group("/comments")
	commentsGroup.Use(middleware.Auth)
	commentsGroup.POST("", commentControllerImpl.PostComment)
	commentsGroup.GET("", commentControllerImpl.GetComment)
	commentsGroup.PUT("/:commentId", commentControllerImpl.UpdateComment)
	commentsGroup.DELETE("/:commentId", commentControllerImpl.DeleteComment)
}

func (controller *CommentControllerImpl) PostComment(c echo.Context) error {
	request := request.Comment{}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.UserId, err = helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	commentResponse, err := controller.Usecase.PostComment(c.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusCreated,
		Message: "Your comment have been successfully posted",
		Data:    commentResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *CommentControllerImpl) GetComment(c echo.Context) error {
	commentResponse, err := controller.Usecase.GetComment(c.Request().Context())
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success get all comments",
		Data:    commentResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *CommentControllerImpl) UpdateComment(c echo.Context) error {
	request := request.Comment{}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	request.Id, err = strconv.Atoi(c.Param("commentId"))
	if err != nil {
		return err
	}

	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.UserId, err = helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	commentResponse, err := controller.Usecase.UpdateComment(c.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success update comment",
		Data:    commentResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *CommentControllerImpl) DeleteComment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		return err
	}

	err = controller.Usecase.DeleteComment(c.Request().Context(), id)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Your comment has been successfully deleted",
	}

	return c.JSON(http.StatusOK, webResponse)
}
