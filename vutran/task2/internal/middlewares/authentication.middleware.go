package middlewares

import (
	"net/http"
	"strings"
	"task2/internal/exceptions"
	services "task2/internal/services/jwt"

	"github.com/gin-gonic/gin"
)

func RequireAuth(jwtMaker services.JWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewUnauthorized("Missing or invalid Authorization header"))
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtMaker.VerifyToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Gắn vào context để handler dùng
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("client_id", claims.ClientID)
		c.Next()
	}
}
