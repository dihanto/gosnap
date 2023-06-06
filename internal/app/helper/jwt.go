package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func GetUserDataFromToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("snapsecret"), nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdFloat, ok := claims["id"].(float64)
		if !ok {
			return 0, fmt.Errorf("failed to retrieve id from token")
		}

		userId := int(userIdFloat)
		return userId, nil
	}

	return 0, fmt.Errorf("invalid token")
}
