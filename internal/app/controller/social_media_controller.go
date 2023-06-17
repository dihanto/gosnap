package controller

import "github.com/labstack/echo/v4"

type SocialMediaController interface {
	PostSocialMedia(ctx echo.Context) error
	GetSocialMedia(ctx echo.Context) error
	UpdateSocialMedia(ctx echo.Context) error
	DeleteSocialMedia(ctx echo.Context) error
}
