package controllers

// import (
// 	"net/http"
// 	"strconv"
// 	"task2/exceptions"
// 	repositories "task2/repositories/user"
// 	services "task2/services/user"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jmoiron/sqlx"
// )

// type UserController struct {
// 	userService *services.UserService
// }

// func NewUserController(db *sqlx.DB) *UserController {
// 	userRepository := repositories.NewUserRepository(db)
// 	userService := services.NewUserService(userRepository)
// 	return &UserController{userService}
// }

// func (userController *UserController) Create() gin.HandlerFunc{
// 	return func(ctx *gin.Context){

// 	}
// }

// func (userController *UserController) FindOne() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

// 		if err != nil {
// 			panic(exceptions.HttpError{
// 				Code: http.StatusUnprocessableEntity,
// 				Message: "Invalid id",
// 			})
// 		}

// 		user, _ := userController.userService.FindOne(ctx.Request.Context(), id)

// 		ctx.JSON(http.StatusOK, user)
// 	}
// }
