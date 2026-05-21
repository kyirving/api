package handler

import (
	"api/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Ping(c *gin.Context) {
	if err := h.svc.Ping(); err != nil {
		Error(c, FAILED, err.Error())
		return
	}
	Success(c, "pong", nil)
}
