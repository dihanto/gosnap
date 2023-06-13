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
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	timeout := viper.GetInt("usecase.timeout")

	e := echo.New()
	e.HTTPErrorHandler = exception.ErrorHandler

	db, _ := config.NewDb()
	validate := validator.New()
	validate.RegisterValidation("email_uniq", helper.ValidateEmailUniq)
	validate.RegisterValidation("username_uniq", helper.ValidateUsernameUniq)

	{
		userRepository := repository.NewUserRepository()
		userUsecase := usecase.NewUserUsecase(userRepository, db, validate, timeout)
		_ = controller.NewUserController(userUsecase, e)
	}

	{
		photoRepository := repository.NewPhotoRepository()
		photoUsecase := usecase.NewPhotoUsecase(photoRepository, db, validate, timeout)
		_ = controller.NewPhotoController(photoUsecase, e)
	}

	{
		commentRepository := repository.NewCommentRepository()
		commentUsecase := usecase.NewCommentUsecase(commentRepository, db, validate, timeout)
		_ = controller.NewCommentController(commentUsecase, e)
	}

	{
		socialMediaRepository := repository.NewSocialMediaRepository()
		socialMediaUsecase := usecase.NewSocialMediaUsecase(socialMediaRepository, db, validate, timeout)
		_ = controller.NewSocialMediaController(socialMediaUsecase, e)
	}

	e.Start(host + ":" + port)

}
