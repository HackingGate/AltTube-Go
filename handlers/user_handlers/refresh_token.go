package user_handlers

import (
	"AltTube-Go/auth"
	"AltTube-Go/database"
	"AltTube-Go/models"
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

	uuid, err := database.ValidateRefreshToken(tokenString)
	if err != nil {
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
	err = database.RemoveRefreshTokenByToken(tokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing old refresh token"})
		return
	}

	user, err := database.GetUserByID(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}

	// Get user_agent from request header
	userAgent := ctx.GetHeader("User-Agent")

	// Get IP address from request
	ipAddress := ctx.ClientIP()

	// Store the new refresh token
	err = database.AddRefreshToken(
		models.RefreshToken{
			Token:     refreshTokenString,
			UserID:    user.ID,
			Expiry:    refreshTokenExpiry,
			UserAgent: userAgent,
			IPAddress: ipAddress,
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing refresh token " + err.Error()})
		return
	}

	refreshToken, err := database.GetRefreshTokenByToken(refreshTokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting refresh token"})
		return
	}

	err = database.AddAccessToken(
		models.AccessToken{
			Token:          accessTokenString,
			UserID:         user.ID,
			Expiry:         accessTokenExpiry,
			RefreshTokenID: refreshToken.ID,
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing access token " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})
}
