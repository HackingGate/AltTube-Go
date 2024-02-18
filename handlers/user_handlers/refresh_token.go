package user_handlers

import (
	"AltTube-Go/auth"
	"AltTube-Go/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// RefreshToken godoc
// @Summary Refresh token
// @Description Refresh token
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} JSON "{"access_token": "access_token", "refresh_token": "refresh_token"}"
// @Security RefreshToken
// @Router /user/refresh_token [post]
func RefreshToken(ctx *gin.Context) {
	// Get authorization header
	tokenString := ctx.GetHeader("Authorization")
	// Assuming token is provided as 'Bearer <token>'
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	uuid, exists := database.ValidateRefreshToken(tokenString)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate new tokens
	accessToken, err := auth.GenerateAccessToken(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating access token"})
		return
	}

	refreshToken, refreshTokenExpiry, err := auth.GenerateRefreshToken(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
		return
	}

	// Replace the old refresh token with the new one
	err = database.RemoveRefreshToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing old refresh token"})
		return
	}
	err = database.AddRefreshToken(refreshToken, uuid, refreshTokenExpiry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing refresh token " + err.Error()})
		return
	}

	auth.AddToken(accessToken)

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
