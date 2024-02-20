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
	authUUIDInterface, exists := ctx.Get("uuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No UUID found in refreshToken"})
		ctx.Abort()
		return
	}

	authUUID, ok := authUUIDInterface.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - UUID format invalid"})
		ctx.Abort()
		return
	}

	// Delete refresh tokens associated with the user
	refreshTokens, err := database.GetAllRefreshTokensByUserID(authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting refresh tokens"})
		return
	}
	for _, refreshToken := range refreshTokens {
		err = database.RemoveAllAccessTokensByRefreshTokenID(refreshToken.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing access tokens associated with refresh refreshToken"})
			return
		}

		err = database.RemoveRefreshTokenByToken(refreshToken.Token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing refreshTokens"})
			return
		}
	}

	// Delete user from database
	err = database.DeleteUserByID(authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	accessTokens, err := database.GetAllAccessTokensByUserID(authUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting access tokens"})
		return
	}

	for _, accessToken := range accessTokens {
		err = database.RemoveAccessToken(accessToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing access refreshToken"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
