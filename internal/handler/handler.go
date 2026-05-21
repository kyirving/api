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

type registerReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, NOT_ACCEPTABLE, err.Error())
		return
	}

	user, err := h.svc.Register(req.Email, req.Password, req.Name)
	if err != nil {
		Error(c, FAILED, err.Error())
		return
	}
	Success(c, "register ok", user)
}

type loginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, NOT_ACCEPTABLE, err.Error())
		return
	}

	token, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		Error(c, NOT_AUTHORIZED, err.Error())
		return
	}
	Success(c, "login ok", gin.H{"token": token})
}
