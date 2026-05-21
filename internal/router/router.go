package router

import (
	"api/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup(userH *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	registerUserRoutes(api, userH)

	return r
}
