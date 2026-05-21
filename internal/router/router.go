package router

import (
	"api/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", h.Ping)
	return r
}
