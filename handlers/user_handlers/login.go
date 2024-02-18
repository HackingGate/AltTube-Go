package user_handlers

import (
	"AltTube-Go/auth"
	"AltTube-Go/database"
	"AltTube-Go/models"
	"AltTube-Go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary Login
// @Description Login
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.Login true "User"
// @Success 200 {string} JSON "{"access_token": "access_token", "refresh_token": "refresh_token"}"
// @Router /user/login [post]s
func Login(ctx *gin.Context) {
	var login models.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := database.GetUserByEmail(login.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	if !utils.CheckPasswordHash(login.Password, foundUser.Password) {
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
	refreshToken, refreshTokenExpiry, err := auth.GenerateRefreshToken(foundUser.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
		return
	}

	// Store refresh token
	err = database.AddRefreshToken(refreshToken, foundUser.UserID, refreshTokenExpiry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing refresh token"})
		return
	}

	// You may also want to add the access token to a list of valid tokens if you're tracking those
	auth.AddToken(accessToken)

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
