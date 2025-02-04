package user_handlers

import (
	"AltTube-Go/auth"
	"AltTube-Go/database"
	"AltTube-Go/dto"
	"AltTube-Go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login
// @Description Login
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body dto.LoginRequest true "User"
// @Success 200 {string} JSON "{"access_token": "access_token", "refresh_token": "refresh_token"}"
// @Router /user/login [post]s
func Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := database.GetUserByEmail(ctx.Request.Context(), loginRequest.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	if !utils.CheckPasswordHash(loginRequest.Password, foundUser.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate access token
	accessTokenString, accessTokenExpiry, err := auth.GenerateAccessToken(foundUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating access token"})
		return
	}

	// Generate refresh token
	refreshTokenString, refreshTokenExpiry, err := auth.GenerateRefreshToken(foundUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
		return
	}

	// Get user_agent from request header
	userAgent := ctx.GetHeader("User-Agent")

	// Get IP address from request
	ipAddress := ctx.ClientIP()

	// Store refresh token
	err = database.AddRefreshToken(ctx.Request.Context(), refreshTokenString, foundUser, refreshTokenExpiry, userAgent, ipAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing refresh token"})
		return
	}

	refreshToken, err := database.GetRefreshTokenByToken(ctx.Request.Context(), refreshTokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting refresh token"})
		return
	}

	// Store access token
	err = database.AddAccessToken(ctx.Request.Context(), accessTokenString, foundUser, accessTokenExpiry, refreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error storing access token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})
}
