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
	e.POST("/users/register", userController.UserRegister)
	e.POST("/users/login", userController.UserLogin)
	e.PUT("/users", userController.UserUpdate, middleware.Auth)
	e.DELETE("/users", userController.UserDelete, middleware.Auth)

	return e
}
