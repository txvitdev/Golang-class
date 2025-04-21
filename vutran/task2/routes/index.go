package routes

import (
	controllers "task2/controllers/user"
	"task2/middlewares"
	repositories "task2/repositories/user"
	services "task2/services/user"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitialRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.ContextTimeout())
	router.Use(middlewares.ErrorHandle())

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	userController := controllers.NewUserController(userService)

	{
		v1 := router.Group("/api/v1")

		v1.POST("/users", userController.Create())
		v1.GET("/users/:id", userController.FindOne())
	}

	return router
}