package user_handlers

import (
	"AltTube-Go/auth"
	"AltTube-Go/database"
	"AltTube-Go/models"
	"AltTube-Go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := database.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	if !utils.CheckPasswordHash(user.Password, foundUser.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate access token
	accessToken, err := auth.GenerateAccessToken(foundUser.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating access token"})
		return
	}

	// Generate refresh token
	refreshToken, err := auth.GenerateRefreshToken(foundUser.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
		return
	}

	// Store refresh token
	auth.AddRefreshToken(refreshToken, foundUser.UserID)

	// You may also want to add the access token to a list of valid tokens if you're tracking those
	auth.AddToken(accessToken)

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
