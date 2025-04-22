package controllers

import (
	"net/http"
	dto "task2/dtos/auth"
	services "task2/services/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService}
}


func (authController *AuthController) SignUp() gin.HandlerFunc{
	return func(ctx *gin.Context){
		var req dto.SignUpDto

		if err:= ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": http.StatusUnprocessableEntity,
				"message": err.Error(),
			})
			return
		}

		user, err := authController.authService.SignUp(ctx, &req)

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