package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecrectKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(id, name, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecrectKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(accessToken string) error {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return jwtSecrectKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
