package user_handlers

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditEmail(ctx *gin.Context) {
	email := ctx.Param("email") // Get email from URL parameter
	newEmail := ctx.PostForm("newEmail")

	err := database.UpdateUserByEmail(email, models.User{Email: newEmail})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating email"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email updated successfully"})
}
