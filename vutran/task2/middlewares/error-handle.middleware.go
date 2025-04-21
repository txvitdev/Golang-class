package middlewares

import (
	"fmt"
	"net/http"
	"task2/exceptions"

	"github.com/gin-gonic/gin"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)

				if err, ok := r.(exceptions.HttpError); ok {
					c.AbortWithStatusJSON(err.Code, exceptions.HttpError{
						Code: err.Code,
						Message: err.Message,
					})
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, exceptions.HttpError{
					Code: http.StatusInternalServerError,
					Message: "Internal server error",
				})
				return
			}
		}()

		c.Next()
	}
}