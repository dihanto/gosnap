package middleware

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Auth = echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte("snapsecret"),
})

func SnapLogger(router *echo.Echo, logFile *os.File) {
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
		Output: logFile,
	}))
}
