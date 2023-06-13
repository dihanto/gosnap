package helper

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GetUserDataFromToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("snapsecret"), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok := claims["id"].(string)
		if !ok {
			return uuid.Nil, err
		}

		userId, err := uuid.Parse(id)
		if err != nil {
			return uuid.Nil, err
		}
		return userId, nil
	}

	return uuid.Nil, err
}
