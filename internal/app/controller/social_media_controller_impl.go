package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/middleware"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/google/uuid"
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
func (socialMediaControllerImpl *SocialMediaControllerImpl) route(e *echo.Echo) {
	socialMediasGroup := e.Group("/socialmedias")
	socialMediasGroup.Use(middleware.Auth)
	socialMediasGroup.POST("", socialMediaControllerImpl.PostSocialMedia)
	socialMediasGroup.GET("", socialMediaControllerImpl.GetSocialMedia)
	socialMediasGroup.PUT("/:socialMediaId", socialMediaControllerImpl.UpdateSocialMedia)
	socialMediasGroup.DELETE("/:socialMediaId", socialMediaControllerImpl.DeleteSocialMedia)
}

func (controller *SocialMediaControllerImpl) PostSocialMedia(c echo.Context) error {

	request := request.SocialMedia{}
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

	socialMediaResponse, err := controller.Usecase.PostSocialMedia(c.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusCreated,
		Message: "Social Media have been successfully posted",
		Data:    socialMediaResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *SocialMediaControllerImpl) GetSocialMedia(c echo.Context) error {
	socialMediaResponse, err := controller.Usecase.GetSocialMedia(c.Request().Context())
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success get all social media",
		Data:    socialMediaResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *SocialMediaControllerImpl) UpdateSocialMedia(c echo.Context) error {

	request := request.SocialMedia{}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	idString := c.Param("socialMediaId")
	request.Id, err = uuid.Parse(idString)
	if err != nil {
		return err
	}

	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.UserId, err = helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	socialMediaResponse, err := controller.Usecase.UpdateSocialMedia(c.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Success update social media",
		Data:    socialMediaResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *SocialMediaControllerImpl) DeleteSocialMedia(c echo.Context) error {
	idString := c.Param("socialMediaId")
	id, err := uuid.Parse(idString)
	if err != nil {
		return err
	}

	err = controller.Usecase.DeleteSocialMedia(c.Request().Context(), id)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status: http.StatusOK,
		Data:   "Your social media has been successfully deleted",
	}

	return c.JSON(http.StatusOK, webResponse)
}
