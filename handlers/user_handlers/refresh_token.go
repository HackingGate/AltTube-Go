package user_handlers

import (
	"AltTube-Go/auth"
	"AltTube-Go/database"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

	uuid, exists := database.ValidateRefreshToken(ctx.Request.Context(), tokenString)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate new tokens
	accessTokenString, accessTokenExpiry, err := auth.GenerateAccessToken(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating access token"})
		return
	}

	refreshTokenString, refreshTokenExpiry, err := auth.GenerateRefreshToken(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
		return
	}

	// Replace the old refresh token with the new one
	err = database.RemoveRefreshTokenByToken(ctx.Request.Context(), tokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing old refresh token"})
		return
	}

	user, err := database.GetUserByID(ctx.Request.Context(), uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}

	// Get user_agent from request header
	userAgent := ctx.GetHeader("User-Agent")

	// Get IP address from request
	ipAddress := ctx.ClientIP()

	// Store the new refresh token
	err = database.AddRefreshToken(ctx.Request.Context(), refreshTokenString, user, refreshTokenExpiry, userAgent, ipAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing refresh token " + err.Error()})
		return
	}

	refreshToken, err := database.GetRefreshTokenByToken(ctx.Request.Context(), refreshTokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting refresh token"})
		return
	}

	err = database.AddAccessToken(ctx.Request.Context(), accessTokenString, user, accessTokenExpiry, refreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing access token " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})
}
