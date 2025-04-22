package middlewares

// import (
// 	"fmt"
// 	"net/http"
// 	"task2/internal/exceptions"
// 	services "task2/internal/services/role"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// func RequirePermission(permission string, roleService services.RoleService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 	  userIDStr, exists := c.Get("user_id")
// 	  if !exists {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewUnauthorized(""))
// 		return
// 	  }

// 	  userID, err := uuid.Parse(fmt.Sprint(userIDStr))
// 	  if err != nil {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, exceptions.NewUnauthorized("Invalid user ID"))
// 		return
// 	  }

// 	  hasPerm, err := roleService.CheckPermission(c.Request.Context(), userID, permission)
// 	  if err != nil {
// 		log.Printf("permission check error: %v", err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "permission check failed"})
// 		return
// 	  }

// 	  if !hasPerm {
// 		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
// 		return
// 	  }

// 	  // Được phép → tiếp tục
// 	  c.Next()
// 	}
//   }
