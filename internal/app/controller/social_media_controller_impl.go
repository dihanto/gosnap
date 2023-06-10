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

type SocialMediaControllerImpl struct {
	Usecase usecase.SocialMediaUsecase
}

func NewSocialMediaController(usecase usecase.SocialMediaUsecase) SocialMediaController {
	return &SocialMediaControllerImpl{
		Usecase: usecase,
	}
}

func (controller *SocialMediaControllerImpl) PostSocialMedia(c echo.Context) error {
	name := c.FormValue("name")
	socialMediaUrl := c.FormValue("socialMediaUrl")
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	userId, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	request := web.PostSocialMedia{
		Name:           name,
		SocialMediaUrl: socialMediaUrl,
		UserId:         userId,
	}

	socialMediaResponse, err := controller.Usecase.PostSocialMedia(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	webResponse := web.WebResponse{
		Status: http.StatusCreated,
		Data:   socialMediaResponse,
	}

	c.Response().Writer.Header().Add("Content-Type", "application/json")
	return c.JSON(http.StatusOK, webResponse)

}

func (controller *SocialMediaControllerImpl) GetSocialMedia(c echo.Context) error {
	socialMediaResponse, err := controller.Usecase.GetSocialMedia(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   socialMediaResponse,
	}

	c.Response().Writer.Header().Add("Content-Type", "application/json")
	return c.JSON(http.StatusOK, webResponse)
}

func (controller *SocialMediaControllerImpl) UpdateSocialMedia(c echo.Context) error {

	idString := c.Param("socialMediaId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	name := c.FormValue("name")
	socialMediaUrl := c.FormValue("socialMediaUrl")

	request := web.UpdateSocialMedia{
		Id:             id,
		Name:           name,
		SocialMediaUrl: socialMediaUrl,
	}

	socialMediaResponse, err := controller.Usecase.UpdateSocialMedia(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   socialMediaResponse,
	}

	c.Response().Writer.Header().Add("Content-Type", "application/json")
	return c.JSON(http.StatusOK, webResponse)
}

func (controller *SocialMediaControllerImpl) DeleteSocialMedia(c echo.Context) error {
	idString := c.Param("socialMediaId")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	err = controller.Usecase.DeleteSocialMedia(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   "Your social media has been succesfully deleted",
	}
	c.Response().Writer.Header().Add("Content-Type", "application/json")
	return c.JSON(http.StatusOK, webResponse)
}
