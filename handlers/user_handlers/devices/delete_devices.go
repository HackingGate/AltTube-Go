package devices

import (
	"AltTube-Go/database"
	"github.com/gin-gonic/gin"
	"net/http"
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
	authUserIDInterface, exists := ctx.Get("uuid")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No UUID found in token"})
		return
	}

	authUserID, ok := authUserIDInterface.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - UUID format invalid"})
		return
	}

	var deleteDevices []uint

	if err := ctx.ShouldBindJSON(&deleteDevices); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshTokens, err := database.GetRefreshTokenByUserID(authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting devices for user"})
		return
	}
	if len(refreshTokens) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No devices can be deleted"})
		return
	}

	// Filter only the refresh tokens that are to be deleted
	// Prepare IDs to be deleted
	var deleteIDs []uint
	for i := range refreshTokens {
		for j := range deleteDevices {
			if refreshTokens[i].ID == deleteDevices[j] {
				deleteIDs = append(deleteIDs, refreshTokens[i].ID)
			}
		}
	}

	// Delete the refresh tokens
	err = database.RemoveRefreshTokensByID(deleteIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting devices"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Devices deleted successfully",
		"deleted": deleteIDs,
	})
}
