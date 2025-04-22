package exceptions

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidatorException(c *gin.Context) {
	if len(c.Errors) > 0 {
		fmt.Println("Here")
		for _, err := range c.Errors {
			var validationErrors validator.ValidationErrors
			if errors.As(err.Err, &validationErrors) {
				c.Status(http.StatusUnprocessableEntity)

				customErrors := make(map[string]string)
				for _, e := range validationErrors {
					field := e.Field()
					switch e.Tag() {
					case "required":
						customErrors[field] = field + " is required"
					case "email":
						customErrors[field] = "Email invalid"
					default:
						customErrors[field] = field + " is invalid"
					}
				}

				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
					"code":     http.StatusUnprocessableEntity,
					"messages": customErrors,
				})
				return
			}
		}
	}
}
