package user_handlers

import (
	"AltTube-Go/database"
	"strings"

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
// @Router /user/logout [post]
func LogoutUser(ctx *gin.Context) {
	// Get authorization header
	tokenString := ctx.GetHeader("Authorization")
	// Assuming currentRefreshToken is provided as 'Bearer <currentRefreshToken>'
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	refreshToken, err := database.GetRefreshTokenByAccessToken(tokenString)
	if err != nil {
		return
	}

	err = database.RemoveAccessToken(tokenString)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error removing access token"})
		return
	}

	err = database.RemoveRefreshTokensByID([]uint{refreshToken.ID})
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error removing refresh token"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Logged out successfully"})
}
