package user_handlers

import (
	"AltTube-Go/database"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
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

	authUUID, ok := authUUIDInterface.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - UserID format invalid"})
		ctx.Abort()
		return
	}

	// Delete refresh tokens associated with the user
	refreshTokens, err := database.GetAllRefreshTokensByUserID(ctx.Request.Context(), authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting refresh tokens"})
		return
	}
	var refreshTokensIDs []uint
	for _, refreshToken := range refreshTokens {
		// Delete access tokens associated with the refresh token
		err = database.RemoveAllAccessTokensByRefreshTokenID(ctx.Request.Context(), refreshToken.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing access tokens associated with refresh refreshToken"})
			return
		}

		refreshTokensIDs = append(refreshTokensIDs, refreshToken.ID)
	}
	err = database.RemoveRefreshTokensByID(ctx.Request.Context(), refreshTokensIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing refresh tokens"})
		return
	}

	// Delete like videos associated with the user
	err = database.RemoveAllLikesByUserID(ctx.Request.Context(), authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing like videos"})
		return
	}

	// Delete user from database
	err = database.DeleteUserByID(ctx.Request.Context(), authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
