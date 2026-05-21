package handler

import (
	"api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

type registerReq struct {
	Username string `json:"username" binding:"required,min=3,max=128"`
	Password string `json:"password" binding:"required,min=6,max=128"`
	Nikename string `json:"nikename" binding:"max=128"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, NOT_ACCEPTABLE, err.Error())
		return
	}

	user, err := h.svc.Register(req.Username, req.Password, req.Nikename)
	if err != nil {
		Error(c, FAILED, err.Error())
		return
	}
	Success(c, "register ok", user)
}

type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, NOT_ACCEPTABLE, err.Error())
		return
	}

	token, err := h.svc.Login(req.Username, req.Password)
	if err != nil {
		Error(c, NOT_AUTHORIZED, err.Error())
		return
	}
	Success(c, "login ok", gin.H{"token": token})
}
