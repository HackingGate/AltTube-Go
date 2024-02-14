package user_handlers

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditEmail(ctx *gin.Context) {
	var editEmail models.EditEmail
	if err := ctx.ShouldBindJSON(&editEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	err := database.UpdateUserByEmail(authEmail, editEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating email"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email updated successfully"})
}
