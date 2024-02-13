package handlers

import (
	"AltTube-Go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Resource(ctx *gin.Context) {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if bearerToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	reqToken := strings.Split(bearerToken, " ")[1]
	for _, token := range auth.GetTokens() {
		if token == reqToken {
			ctx.JSON(http.StatusOK, gin.H{
				"data": "resource data",
			})
			return
		}
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
}
