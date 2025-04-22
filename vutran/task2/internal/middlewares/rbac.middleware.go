package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"task2/internal/exceptions"
	services "task2/internal/services/role"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequireRole(role string, roleService services.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewUnauthorized(""))
			return
		}

		_, err := uuid.Parse(fmt.Sprint(userIDStr))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewUnauthorized("Invalid user ID"))
			return
		}

		hasPerm, httpError := roleService.CheckRoleOfUser(c.Request.Context(), userIDStr.(string), role)
		if httpError.Code != 0 {
			log.Printf("permission check error: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, exceptions.NewInternal("permission check failed"))
			return
		}

		if !hasPerm {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewForbidden(""))
			return
		}
		c.Next()
	}
}
