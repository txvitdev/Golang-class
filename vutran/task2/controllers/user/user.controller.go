package controllers

import (
	"net/http"
	"strconv"
	dto "task2/dtos/auth"
	"task2/exceptions"
	services "task2/services/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService}
}


func (userController *UserController) Create() gin.HandlerFunc{
	return func(ctx *gin.Context){
		var req dto.SignUpDto

		if err:= ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": http.StatusUnprocessableEntity,
				"message": err.Error(),
			})
			return
		}

		user, err := userController.userService.Save(ctx, req.Email, req.Password)

		if err.Code != 0 {
			ctx.JSON(err.Code, gin.H{
				"code": err.Code,
				"message": err.Message,
			})

			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (userController *UserController) FindOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			panic(exceptions.HttpError{
				Code: http.StatusUnprocessableEntity,
				Message: "Invalid id",
			})
		}

		user, _ := userController.userService.FindOne(ctx.Request.Context(), id)

		ctx.JSON(http.StatusOK, user)
	}
}