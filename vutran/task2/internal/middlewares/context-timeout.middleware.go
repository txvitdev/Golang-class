package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
)

func ContextTimeout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(ctx.Request.Context(), 3 * 1000)
		defer cancel()

		ctx.Next()
	}
}