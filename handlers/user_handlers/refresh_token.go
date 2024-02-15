package user_handlers

import (
	"AltTube-Go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RefreshToken(ctx *gin.Context) {
	// Get authorization header
	tokenString := ctx.GetHeader("Authorization")
	// Assuming token is provided as 'Bearer <token>'
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	uuid, exists := auth.ValidateRefreshToken(tokenString)
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

	refreshToken, err := auth.GenerateRefreshToken(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
		return
	}

	// Replace the old refresh token with the new one
	auth.RemoveRefreshToken(tokenString)
	auth.AddRefreshToken(refreshToken, uuid)

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
