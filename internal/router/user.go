package router

import (
	"api/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler) {
	rg.POST("/register", h.Register)
	rg.POST("/login", h.Login)
}
