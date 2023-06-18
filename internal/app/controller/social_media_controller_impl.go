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

type SocialMediaControllerImpl struct {
	Usecase usecase.SocialMediaUsecase
	Route   *echo.Echo
}

func NewSocialMediaController(usecase usecase.SocialMediaUsecase, route *echo.Echo) SocialMediaController {
	controller := &SocialMediaControllerImpl{
		Usecase: usecase,
		Route:   route,
	}
	controller.route(route)
	return controller
}
func (socialMediaControllerImpl *SocialMediaControllerImpl) route(echo *echo.Echo) {
	socialMediasGroup := echo.Group("/socialmedias")
	socialMediasGroup.Use(middleware.Auth)
	socialMediasGroup.POST("", socialMediaControllerImpl.PostSocialMedia)
	socialMediasGroup.GET("", socialMediaControllerImpl.GetSocialMedia)
	socialMediasGroup.PUT("/:socialMediaId", socialMediaControllerImpl.UpdateSocialMedia)
	socialMediasGroup.DELETE("/:socialMediaId", socialMediaControllerImpl.DeleteSocialMedia)
}

func (controller *SocialMediaControllerImpl) PostSocialMedia(ctx echo.Context) error {

	request := request.SocialMedia{}
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

	socialMediaResponse, err := controller.Usecase.PostSocialMedia(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusCreated,
		Message: "Social Media have been successfully posted",
		Data:    socialMediaResponse,
	}

	return ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *SocialMediaControllerImpl) GetSocialMedia(ctx echo.Context) error {
	socialMediaResponse, err := controller.Usecase.GetSocialMedia(ctx.Request().Context())
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success get all social media",
		Data:    socialMediaResponse,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}

func (controller *SocialMediaControllerImpl) UpdateSocialMedia(ctx echo.Context) error {

	request := request.SocialMedia{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	request.Id, err = strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		return err
	}

	authHeader := ctx.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.UserId, err = helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	socialMediaResponse, err := controller.Usecase.UpdateSocialMedia(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success update social media",
		Data:    socialMediaResponse,
	}

	return ctx.JSON(http.StatusOK, webResponse)
}

func (controller *SocialMediaControllerImpl) DeleteSocialMedia(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		return err
	}

	err = controller.Usecase.DeleteSocialMedia(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status: http.StatusOK,
		Data:   "Your social media has been successfully deleted",
	}

	return ctx.JSON(http.StatusOK, webResponse)
}
