package middlewares

import (
	"net/http"
	"task2/exceptions"

	"github.com/gin-gonic/gin"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, exceptions.HttpError{
					Code: http.StatusInternalServerError,
					Message: "Internal server error",
				})
			}
		}()

		c.Next()

	}
}