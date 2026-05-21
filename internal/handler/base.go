package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS        = 0
	FAILED         = 1
	NOT_AUTHORIZED = 401
	NOT_FOUND      = 404
	NOT_ACCEPTABLE = 406
	INTERNAL_ERROR = 500
	NETWORK_ERROR  = 502
)

var respMsg = map[int]string{
	SUCCESS:        "ok",
	FAILED:         "failed",
	NOT_AUTHORIZED: "not authorized",
	NOT_FOUND:      "not found",
	NOT_ACCEPTABLE: "not acceptable",
	INTERNAL_ERROR: "internal error",
	NETWORK_ERROR:  "network error",
}

var bizToHTTP = map[int]int{
	SUCCESS:        http.StatusOK,
	FAILED:         http.StatusInternalServerError,
	NOT_AUTHORIZED: http.StatusUnauthorized,
	NOT_FOUND:      http.StatusNotFound,
	NOT_ACCEPTABLE: http.StatusNotAcceptable,
	INTERNAL_ERROR: http.StatusInternalServerError,
	NETWORK_ERROR:  http.StatusBadGateway,
}

func Success(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    SUCCESS,
		"message": msg,
		"data":    data,
	})
}

func Error(c *gin.Context, bizCode int, msg string) {
	httpStatus, ok := bizToHTTP[bizCode]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}
	c.JSON(httpStatus, gin.H{
		"code":    bizCode,
		"message": respMsg[bizCode] + " " + msg,
	})
}
