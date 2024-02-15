package user_handlers

import (
	"AltTube-Go/auth"
	"AltTube-Go/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(ctx *gin.Context) {
	authUUIDInterface, exists := ctx.Get("uuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No UUID found in token"})
		ctx.Abort()
		return
	}

	authUUID, ok := authUUIDInterface.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - UUID format invalid"})
		ctx.Abort()
		return
	}

	// Delete user from database
	err := database.DeleteUserByUserID(authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	auth.RemoveToken(ctx.GetHeader("Authorization"))

	refreshTokens, err := database.GetAllRefreshTokensByUserID(authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting refresh tokens"})
		return
	}

	for _, token := range refreshTokens {
		err = database.RemoveRefreshToken(token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing refresh token"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
