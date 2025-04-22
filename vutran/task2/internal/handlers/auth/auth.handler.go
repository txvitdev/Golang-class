package handlers

import (
	"fmt"
	"net/http"
	dto "task2/internal/dtos/auth"
	services "task2/internal/services/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(AuthService services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService}
}

func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/auth")
	group.POST("/login", h.Login())
	group.POST("/register", h.Register())
}

func (h *AuthHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.SignInDto
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		token, err := h.authService.Login(c.Request.Context(), req.Email, req.Password, req.ClientID)
		if err.Code != 0 {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, token)
		return
	}
}

func (h *AuthHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.SignUpDto
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		err := h.authService.Register(c.Request.Context(), req.Email, req.Password, req.FullName)
		if err.Code != 0 {
			fmt.Println(err)
			c.Error(err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": "user registered"})
		return
	}
}
