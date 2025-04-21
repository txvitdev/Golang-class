package routes

import (
	controllers "task2/controllers/user"
	dto "task2/dtos/auth"
	"task2/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitialRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.ContextTimeout())
	router.Use(middlewares.ErrorHandle())

	userController := controllers.NewUserController(db)

	{
		v1 := router.Group("/api/v1")
		userRouter := v1.Group("/users")
		authRouter := v1.Group("/auth")
		{
			// Public api
			authRouter.POST("/sign-up", gin.Bind(dto.SignUpDto{}), userController.Create())
		}
		{
			userRouter.POST("/", gin.Bind(dto.SignUpDto{}), userController.Create())

			userRouter.Use(middlewares.AuthenticationMiddleware())

			userRouter.GET("/:id", userController.FindOne())
		}

	}

	return router
}
