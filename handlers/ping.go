package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
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
