package devices

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDevices godoc
// @Summary Get devices
// @Description Get devices
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Device "GetDevices"
// @Security AccessToken
// @Router /user/devices [get]
func GetDevices(ctx *gin.Context) {
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

	refreshTokens, err := database.GetRefreshTokenByUserID(authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting devices for user"})
		return
	}
	if len(refreshTokens) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No devices found"})
		return
	}
	// Filter only the necessary fields
	var devices []models.Device
	for i := range refreshTokens {
		devices = append(devices, models.Device{
			ID:         refreshTokens[i].ID,
			LastActive: refreshTokens[i].CreatedAt,
			UserAgent:  refreshTokens[i].UserAgent,
			IPAddress:  refreshTokens[i].IPAddress,
		})
	}

	ctx.JSON(http.StatusOK, devices)
}
