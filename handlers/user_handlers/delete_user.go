package user_handlers

import (
	"AltTube-Go/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} JSON "{"message": "User deleted successfully"}"
// @Security AccessToken
// @Router /user/ [delete]
func DeleteUser(ctx *gin.Context) {
	authUUIDInterface, exists := ctx.Get("UserID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No UserID found in refreshToken"})
		ctx.Abort()
		return
	}

	authUUID, ok := authUUIDInterface.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - UserID format invalid"})
		ctx.Abort()
		return
	}

	// Delete refresh tokens associated with the user
	refreshTokens, err := database.GetAllRefreshTokensByUserID(authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting refresh tokens"})
		return
	}
	var refreshTokensIDs []uint
	for _, refreshToken := range refreshTokens {
		// Delete access tokens associated with the refresh token
		err = database.RemoveAllAccessTokensByRefreshTokenID(refreshToken.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing access tokens associated with refresh refreshToken"})
			return
		}

		refreshTokensIDs = append(refreshTokensIDs, refreshToken.ID)
	}

	err = database.RemoveRefreshTokensByID(refreshTokensIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing refresh tokens"})
		return
	}

	// Delete user from database
	err = database.DeleteUserByID(authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
