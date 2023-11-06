package main

import (
	"log"

	"github.com/dihanto/gosnap/internal/app/config"
	"github.com/dihanto/gosnap/internal/app/controller"
	"github.com/dihanto/gosnap/internal/app/exception"
	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/middleware"
	"github.com/dihanto/gosnap/internal/app/repository"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {

	config.InitLoadConfiguration()
	serverHost := viper.GetString("server.host")
	serverPort := viper.GetString("server.port")
	usecaseTimeout := viper.GetInt("usecase.timeout")

	router := echo.New()
	router.HTTPErrorHandler = exception.ErrorHandler
	router.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	logFile := config.InitLogFile()
	defer logFile.Close()
	middleware.SnapLogger(router, logFile)

	databaseConnection, err := config.InitDatabaseConnection()
	if err != nil {
		log.Fatalln(err)
	}
	validate := validator.New()
	validate.RegisterValidation("email_uniq", helper.ValidateEmailUniq)
	validate.RegisterValidation("username_uniq", helper.ValidateUsernameUniq)
	validate.RegisterValidation("likes", helper.ValidateOneUserOneLike)
	validate.RegisterValidation("follow", helper.ValidateUserNotFollowTwice)

	{
		userRepository := repository.NewUserRepository(databaseConnection)
		userUsecase := usecase.NewUserUsecase(userRepository, validate, usecaseTimeout)
		controller.NewUserController(userUsecase, router)
	}

	{
		photoRepository := repository.NewPhotoRepository(databaseConnection)
		photoUsecase := usecase.NewPhotoUsecase(photoRepository, validate, usecaseTimeout)
		controller.NewPhotoController(photoUsecase, router)
	}

	{
		commentRepository := repository.NewCommentRepository(databaseConnection)
		commentUsecase := usecase.NewCommentUsecase(commentRepository, validate, usecaseTimeout)
		controller.NewCommentController(commentUsecase, router)
	}

	{
		socialMediaRepository := repository.NewSocialMediaRepository(databaseConnection)
		socialMediaUsecase := usecase.NewSocialMediaUsecase(socialMediaRepository, validate, usecaseTimeout)
		controller.NewSocialMediaController(socialMediaUsecase, router)
	}

	{
		followRepository := repository.NewFollowRepositoryImpl(databaseConnection)
		followUsecase := usecase.NewFollowUsecaseImpl(followRepository, validate, usecaseTimeout)
		controller.NewFollowControllerImpl(followUsecase, router)
	}
	{
		likeRepository := repository.NewLikeRepository(databaseConnection)
		likeUsecase := usecase.NewLikeUsecaseImpl(likeRepository, validate, usecaseTimeout)
		controller.NewLikeController(likeUsecase, router)
	}

	err = router.Start(serverHost + ":" + serverPort)
	if err != nil {
		log.Fatalln(err)
	}

}
