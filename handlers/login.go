package handlers

import (
	"AltTube-Go/auth"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	token, _ := auth.GenerateJWT()
	auth.AddToken(token)

	ctx.JSON(200, gin.H{
		"token": token,
	})
}
