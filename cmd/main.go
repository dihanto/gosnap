package main

import (
	"github.com/dihanto/gosnap/internal/app/config"
	"github.com/dihanto/gosnap/internal/app/controller"
	"github.com/dihanto/gosnap/internal/app/exception"
	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/repository"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {

	config.InitLoadConfiguration()
	serverHost := viper.GetString("server.host")
	serverPort := viper.GetString("server.port")
	usecaseTimeout := viper.GetInt("usecase.timeout")

	echo := echo.New()
	echo.HTTPErrorHandler = exception.ErrorHandler

	databaseConnection, _ := config.InitDatabaseConnection()
	validate := validator.New()
	validate.RegisterValidation("email_uniq", helper.ValidateEmailUniq)
	validate.RegisterValidation("username_uniq", helper.ValidateUsernameUniq)

	{
		userRepository := repository.NewUserRepository()
		userUsecase := usecase.NewUserUsecase(userRepository, databaseConnection, validate, usecaseTimeout)
		_ = controller.NewUserController(userUsecase, echo)
	}

	{
		photoRepository := repository.NewPhotoRepository()
		photoUsecase := usecase.NewPhotoUsecase(photoRepository, databaseConnection, validate, usecaseTimeout)
		_ = controller.NewPhotoController(photoUsecase, echo)
	}

	{
		commentRepository := repository.NewCommentRepository()
		commentUsecase := usecase.NewCommentUsecase(commentRepository, databaseConnection, validate, usecaseTimeout)
		_ = controller.NewCommentController(commentUsecase, echo)
	}

	{
		socialMediaRepository := repository.NewSocialMediaRepository()
		socialMediaUsecase := usecase.NewSocialMediaUsecase(socialMediaRepository, databaseConnection, validate, usecaseTimeout)
		_ = controller.NewSocialMediaController(socialMediaUsecase, echo)
	}

	echo.Start(serverHost + ":" + serverPort)

}
