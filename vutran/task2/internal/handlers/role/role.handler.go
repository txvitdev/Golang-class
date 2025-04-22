package handlers

import (
	"task2/internal/middlewares"
	jwtService "task2/internal/services/jwt"
	roleService "task2/internal/services/role"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleService roleService.RoleService
	jwtService  jwtService.JWTMaker
}

func NewRoleHandler(roleService roleService.RoleService, jwtService jwtService.JWTMaker) *RoleHandler {
	return &RoleHandler{roleService, jwtService}
}

func (h *RoleHandler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/roles")
	group.Use(middlewares.RequireAuth(h.jwtService))
	group.Use(middlewares.RequireRole("admin", h.roleService))
}
