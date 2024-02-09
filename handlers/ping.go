package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create ping handler
// @Summary ping
// @Description ping
// @Tags ping
// @Accept  json
// @Produce  json
// @Success 200 {string} JSON "{"message": "pong"}"
// @Router /ping [get]

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
