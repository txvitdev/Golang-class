package middlewares

import (
	"net/http"
	"task2/internal/exceptions"

	"github.com/gin-gonic/gin"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(exceptions.HttpError); ok {
					c.AbortWithStatusJSON(err.Code, exceptions.HttpError{
						Code:    err.Code,
						Message: err.Message,
					})

					return
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, exceptions.NewInternal(""))
				return
			}

			exceptions.HandleValidatorException(c)
			return
		}()

		c.Next()
	}
}
