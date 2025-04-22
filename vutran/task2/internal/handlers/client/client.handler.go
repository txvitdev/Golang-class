package handlers

import (
	"net/http"
	dto "task2/internal/dtos/client"
	services "task2/internal/services/client"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	clientService services.ClientService
}

func NewClientHandler(clientService services.ClientService) *ClientHandler {
	return &ClientHandler{clientService}
}

func (h *ClientHandler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/clients")
	group.POST("/", h.RegisterClient)
}

func (h *ClientHandler) RegisterClient(c *gin.Context) {
	var req dto.RegisterClientRequestDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	client, err := h.clientService.RegisterClient(c.Request.Context(), req.Name, req.RedirectURIs, req.AdminEmail,
		req.AdminPassword,
		req.AdminFullName)
	if err.Code != 0 {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"client_id":     client.ClientID,
		"client_secret": client.ClientSecret,
	})
}
