package user_handlers

import (
	"AltTube-Go/database"
	"AltTube-Go/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EditEmail godoc
// @Summary Edit email
// @Description Edit email
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body dto.UpdateEmailRequest true "User"
// @Success 200 {string} JSON "{"message": "Email updated successfully"}"
// @Security AccessToken
// @Router /user/email [patch]
func EditEmail(ctx *gin.Context) {
	var updateEmailRequest dto.UpdateEmailRequest
	if err := ctx.ShouldBindJSON(&updateEmailRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authUserIDInterface, exists := ctx.Get("UserID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No UserID found in token"})
		return
	}

	authUserID, ok := authUserIDInterface.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - UserID format invalid"})
		return
	}

	err := database.UpdateUserByID(ctx.Request.Context(), authUserID, updateEmailRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating email"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email updated successfully"})
}
