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
	refreshTokenString := ctx.GetHeader("Authorization")
	// Assuming token is provided as 'Bearer <token>'
	refreshTokenString = strings.TrimPrefix(refreshTokenString, "Bearer ")

	exists, err := database.ValidateRefreshToken(ctx.Request.Context(), refreshTokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error validating refresh token"})
		return
	}
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Find user associated with the refresh token
	user, err := database.GetUserByRefreshToken(ctx.Request.Context(), refreshTokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}

	// Generate new access token
	newAccessTokenString, accessTokenExpiry, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating access token"})
		return
	}

	// Generate new refresh token
	newRefreshTokenString, refreshTokenExpiry, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
		return
	}

	// Remove access tokens associated with the old refresh token
	err = database.RemoveAccessTokenByRefreshToken(ctx.Request.Context(), refreshTokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing access token"})
		return
	}

	// Remove the old refresh token
	err = database.RemoveRefreshTokenByToken(ctx.Request.Context(), refreshTokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing old refresh token"})
		return
	}

	// Get user_agent from request header
	userAgent := ctx.GetHeader("User-Agent")

	// Get IP address from request
	ipAddress := ctx.ClientIP()

	// Store the new refresh token
	err = database.AddRefreshToken(ctx.Request.Context(), newRefreshTokenString, user, refreshTokenExpiry, userAgent, ipAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing refresh token " + err.Error()})
		return
	}

	refreshToken, err := database.GetRefreshTokenByToken(ctx.Request.Context(), newRefreshTokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting refresh token"})
		return
	}

	err = database.AddAccessToken(ctx.Request.Context(), newAccessTokenString, user, accessTokenExpiry, refreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing access token " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessTokenString,
		"refresh_token": newRefreshTokenString,
	})
}
