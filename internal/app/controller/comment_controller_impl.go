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
func (commentControllerImpl *CommentControllerImpl) route(echo *echo.Echo) {
	commentsGroup := echo.Group("/comments")
	commentsGroup.Use(middleware.Auth)
	commentsGroup.POST("", commentControllerImpl.PostComment)
	commentsGroup.GET("", commentControllerImpl.GetComment)
	commentsGroup.PUT("/:commentId", commentControllerImpl.UpdateComment)
	commentsGroup.DELETE("/:commentId", commentControllerImpl.DeleteComment)
}

func (controller *CommentControllerImpl) PostComment(ctx echo.Context) error {
	request := request.Comment{}
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

	commentResponse, err := controller.Usecase.PostComment(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusCreated,
		Message: "Your comment have been successfully posted",
		Data:    commentResponse,
	}

	return ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *CommentControllerImpl) GetComment(ctx echo.Context) error {
	commentResponse, err := controller.Usecase.GetComment(ctx.Request().Context())
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success get all comments",
		Data:    commentResponse,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CommentControllerImpl) UpdateComment(ctx echo.Context) error {
	request := request.Comment{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	request.Id, err = strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		return err
	}

	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.UserId, err = helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	commentResponse, err := controller.Usecase.UpdateComment(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success update comment",
		Data:    commentResponse,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CommentControllerImpl) DeleteComment(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		return err
	}

	err = controller.Usecase.DeleteComment(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Your comment has been successfully deleted",
	}

	return ctx.JSON(http.StatusOK, webResponse)
}
