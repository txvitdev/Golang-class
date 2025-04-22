package services

import (
	"errors"
	"task2/internal/config"
	dto "task2/internal/dtos/auth"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker interface {
	CreateToken(userID, email, clientID string) (string, error)
	VerifyToken(tokenStr string) (*dto.TokenClaims, error)
}

type jwtMaker struct {
	secret string
}

func NewJWTMaker(secret string) JWTMaker {
	return &jwtMaker{secret}
}

func (m *jwtMaker) CreateToken(userID, email, clientID string) (string, error) {
	claims := jwt.MapClaims{
		"sub":       userID,
		"email":     email,
		"client_id": clientID,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secret))
}

func (m *jwtMaker) VerifyToken(tokenStr string) (*dto.TokenClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(m.secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return &dto.TokenClaims{
		UserID:   claims["sub"].(string),
		Email:    claims["email"].(string),
		ClientID: claims["client_id"].(string),
	}, nil
}

func ProvideJWTMaker(cfg *config.Config) JWTMaker {
	return NewJWTMaker(cfg.JwtSecret)
}
