package user_handlers

import (
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
	err := database.DeleteUserByUUID(authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
