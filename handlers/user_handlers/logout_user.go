package user_handlers

import (
	"AltTube-Go/auth"
	"github.com/gin-gonic/gin"
)

// LogoutUser godoc
// @Summary Logout user
// @Description Logout user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} JSON "{"message": "Logged out successfully"}"
// @Security AccessToken
// @Router /user/logout [get]
func LogoutUser(ctx *gin.Context) {
	auth.RemoveToken(ctx.GetHeader("Authorization"))
	ctx.JSON(200, gin.H{"message": "Logged out successfully"})
}
