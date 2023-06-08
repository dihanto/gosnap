package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/dihanto/gosnap/internal/app/helper"
	"github.com/dihanto/gosnap/internal/app/usecase"
	"github.com/dihanto/gosnap/model/web"
	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &UserControllerImpl{
		UserUsecase: userUsecase,
	}
}
func (controller *UserControllerImpl) UserRegister(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	ageString := c.FormValue("age")
	age, err := strconv.Atoi(ageString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	request := web.UserRegister{
		Username: username,
		Email:    email,
		Password: password,
		Age:      age,
	}

	userResponse, err := controller.UserUsecase.UserRegister(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   userResponse,
	}

	c.Response().Writer.Header().Add("Content-Type", "application/json")
	return c.JSON(http.StatusOK, webResponse)

}

func (controller *UserControllerImpl) UserLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, id, err := controller.UserUsecase.UserLogin(c.Request().Context(), username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["id"] = id
	claims["level"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("snapsecret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   t,
	}

	return c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) UserUpdate(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	id, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	username := c.FormValue("username")
	email := c.FormValue("email")
	ageString := c.FormValue("age")
	age, err := strconv.Atoi(ageString)
	if err != nil {
		panic(err)
	}

	request := web.UserUpdate{
		Id:       id,
		Username: username,
		Email:    email,
		Age:      age,
	}

	userResponse, err := controller.UserUsecase.UserUpdate(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   userResponse,
	}

	c.Response().Writer.Header().Add("Content-Type", "application/json")
	return c.JSON(http.StatusOK, webResponse)

}

func (controller *UserControllerImpl) UserDelete(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	id, err := helper.GetUserDataFromToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	controller.UserUsecase.UserDelete(c.Request().Context(), id)

	webResponse := web.WebResponse{
		Status: http.StatusOK,
		Data:   "user deleted",
	}

	c.Response().Writer.Header().Add("Content-Type", "application/json")
	return c.JSON(http.StatusOK, webResponse)

}
