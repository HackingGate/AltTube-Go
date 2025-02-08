package devices

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/hackinggate/alttube-go/database"
	"github.com/hackinggate/alttube-go/ent"

	"github.com/gin-gonic/gin"
)

// DeleteDevices godoc
// @Summary Delete devices
// @Description Delete devices
// @Tags user
// @Accept  json
// @Produce  json
// @Param devices body []uint true "Devices to be deleted"
// @Success 200 {string} JSON "{"message": "Devices deleted successfully", "deleted": [1, 2, 3]}"
// @Security AccessToken
// @Router /user/devices [delete]
func DeleteDevices(ctx *gin.Context) {
	authUserIDInterface, exists := ctx.Get("UserID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No UserID found in token"})
		return
	}

	authUserID, ok := authUserIDInterface.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - UserID format invalid"})
		return
	}

	var deleteDevices []uint

	if err := ctx.ShouldBindJSON(&deleteDevices); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete refresh tokens associated with the user
	refreshTokens, err := database.GetAllRefreshTokensByUserID(ctx.Request.Context(), authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting devices for user"})
		return
	}
	// Filter refreshTokens with deleteDevices
	filteredRefreshTokens := make([]*ent.RefreshToken, 0)
	for _, refreshToken := range refreshTokens {
		for _, device := range deleteDevices {
			if refreshToken.ID == device {
				filteredRefreshTokens = append(filteredRefreshTokens, refreshToken)
				break
			}
		}
	}
	refreshTokens = filteredRefreshTokens
	if len(refreshTokens) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No devices can be deleted"})
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

	// Delete the refresh tokens
	err = database.RemoveRefreshTokensByID(ctx.Request.Context(), refreshTokensIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting devices"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Devices deleted successfully",
		"deleted": refreshTokensIDs,
	})
}
