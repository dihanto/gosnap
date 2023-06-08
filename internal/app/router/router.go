package router

import (
	"github.com/dihanto/gosnap/internal/app/config"
	"github.com/dihanto/gosnap/internal/app/controller"
	"github.com/dihanto/gosnap/internal/app/middleware"
	"github.com/dihanto/gosnap/internal/app/repository"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	db := config.NewDb()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository, db, validate)
	userController := controller.NewUserController(userUsecase)

	photoRepository := repository.NewPhotoRepository()
	photoUsecase := usecase.NewPhotoUsecase(photoRepository, db, validate)
	photoController := controller.NewPhotoController(photoUsecase)

	commentRepository := repository.NewCommentRepository()
	commentUsecase := usecase.NewCommentUsecase(commentRepository, db, validate)
	commentController := controller.NewCommentController(commentUsecase)

	e.POST("/users/register", userController.UserRegister)
	e.POST("/users/login", userController.UserLogin)
	e.PUT("/users", userController.UserUpdate, middleware.Auth)
	e.DELETE("/users", userController.UserDelete, middleware.Auth)

	e.POST("/photos", photoController.PostPhoto, middleware.Auth)
	e.GET("/photos", photoController.GetPhoto, middleware.Auth)
	e.PUT("/photos/:photoId", photoController.UpdatePhoto, middleware.Auth)
	e.DELETE("/photos/:photoId", photoController.DeletePhoto, middleware.Auth)

	e.POST("/comments", commentController.PostComment, middleware.Auth)
	e.GET("/comments", commentController.GetComment, middleware.Auth)
	e.PUT("/comments/:commentId", commentController.UpdateComment, middleware.Auth)
	e.DELETE("/comments/:commentId", commentController.DeleteComment, middleware.Auth)

	return e
}
