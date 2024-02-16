package user_handlers

import (
	"AltTube-Go/auth"
	"github.com/gin-gonic/gin"
)

func LogoutUser(ctx *gin.Context) {
	auth.RemoveToken(ctx.GetHeader("Authorization"))
	ctx.JSON(200, gin.H{"message": "Logged out successfully"})
}
