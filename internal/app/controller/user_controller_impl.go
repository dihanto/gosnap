package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/middleware"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/dihanto/gosnap/model/web/request"
	"github.com/dihanto/gosnap/model/web/response"
	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserUsecase usecase.UserUsecase
	Route       *echo.Echo
}

func NewUserController(userUsecase usecase.UserUsecase, route *echo.Echo) UserController {
	controller := &UserControllerImpl{
		UserUsecase: userUsecase,
		Route:       route,
	}
	controller.route(route)
	return controller
}
func (controller *UserControllerImpl) route(e *echo.Echo) {
	usersGroup := e.Group("/users")
	usersGroup.POST("/register", controller.UserRegister)
	usersGroup.POST("/login", controller.UserLogin)
	usersGroup.Use(middleware.Auth)
	usersGroup.PUT("", controller.UserUpdate)
	usersGroup.DELETE("", controller.UserDelete)
}

func (controller *UserControllerImpl) UserRegister(c echo.Context) error {
	request := request.UserRegister{}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	userResponse, err := controller.UserUsecase.UserRegister(c.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusCreated,
		Message: "You have been successfully registered",
		Data:    userResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) UserLogin(c echo.Context) error {
	request := request.UserLogin{}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	username := request.Username
	password := request.Password

	res, id, err := controller.UserUsecase.UserLogin(c.Request().Context(), username, password)
	if err != nil {
		return err
	}

	if !res {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["id"] = id
	claims["level"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("snapsecret"))
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Login Success",
		Data:    t,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) UserUpdate(c echo.Context) error {

	request := request.UserUpdate{}
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	request.Id, err = helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	userResponse, err := controller.UserUsecase.UserUpdate(c.Request().Context(), request)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "User update success",
		Data:    userResponse,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) UserDelete(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	id, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return err
	}

	err = controller.UserUsecase.UserDelete(c.Request().Context(), id)
	if err != nil {
		return err
	}

	webResponse := response.WebResponse{
		Status:  http.StatusOK,
		Message: "Your account has been successfully deleted",
	}

	return c.JSON(http.StatusOK, webResponse)
}
