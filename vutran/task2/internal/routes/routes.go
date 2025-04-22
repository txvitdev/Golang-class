package routes

import (
	authHandler "task2/internal/handlers/auth"
	clientHandler "task2/internal/handlers/client"
	"task2/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(clientHandler *clientHandler.ClientHandler, authHandler *authHandler.AuthHandler) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.ErrorHandle())
	r.Use(middlewares.ContextTimeout())

	api := r.Group("/api/v1")
	clientHandler.RegisterRoutes(api)
	authHandler.RegisterRoutes(api)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
