package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
)

var Auth = echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte("snapsecret"),
})
