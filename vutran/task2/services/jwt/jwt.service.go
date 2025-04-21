package services

import (
	"task2/config"
	dto "task2/dtos/auth"
	"task2/exceptions"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(payload dto.JwtPayloadDto) (string, error) {
	config := config.SetupConfig()
	claims := jwt.MapClaims{
		"sub":     payload.Sub,
		"options": payload.Options,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(config.JwtSecret)

	if err != nil {
		return "", exceptions.NewInternal("")
	}

	return tokenString, nil
}
