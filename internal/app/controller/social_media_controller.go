package controller

import "github.com/labstack/echo/v4"

type SocialMediaController interface {
	PostSocialMedia(c echo.Context) error
	GetSocialMedia(c echo.Context) error
	UpdateSocialMedia(c echo.Context) error
	DeleteSocialMedia(c echo.Context) error
}
