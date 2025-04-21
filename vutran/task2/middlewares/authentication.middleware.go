package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"task2/config"
	"task2/exceptions"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		config := config.SetupConfig()
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewUnauthorized("Missing token"))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodES256.Alg()}))

		if err != nil || !token.Valid {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewUnauthorized("Token invalid"))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("id", claims["sub"])
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewUnauthorized("Token invalid"))
			return
		}

		ctx.Next()
	}
}
