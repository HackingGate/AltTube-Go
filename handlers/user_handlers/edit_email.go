package user_handlers

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditEmail(ctx *gin.Context) {
	authEmailInterface, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No email found in token"})
		return
	}

	authEmail, ok := authEmailInterface.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - Email format invalid"})
		return
	}

	newEmail := ctx.PostForm("newEmail")

	err := database.UpdateUserByEmail(authEmail, models.User{Email: newEmail})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating email"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email updated successfully"})
}
