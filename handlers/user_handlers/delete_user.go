package user_handlers

import (
	"AltTube-Go/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(ctx *gin.Context) {
	authEmailInterface, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No email found in token"})
		ctx.Abort()
		return
	}

	authEmail, ok := authEmailInterface.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - Email format invalid"})
		ctx.Abort()
		return
	}

	// Delete user from database
	err := database.DeleteUserByEmail(authEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
