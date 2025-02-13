package devices

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/hackinggate/alttube-go/database"
	"github.com/hackinggate/alttube-go/models"

	"github.com/gin-gonic/gin"
)

// GetDevices godoc
// @Summary Get devices
// @Description Get devices
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} models.DeviceListResponse
// @Security AccessToken
// @Router /user/devices [get]
func GetDevices(ctx *gin.Context) {
	// Get authorization header
	tokenString := ctx.GetHeader("Authorization")
	// Assuming currentRefreshToken is provided as 'Bearer <currentRefreshToken>'
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	authUserIDInterface, exists := ctx.Get("UserID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No UserID found in currentRefreshToken"})
		return
	}

	authUserID, ok := authUserIDInterface.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - UserID format invalid"})
		return
	}

	refreshTokens, err := database.GetAllRefreshTokensByUserID(ctx.Request.Context(), authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting devices for user"})
		return
	}
	if len(refreshTokens) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No devices found"})
		return
	}
	// Filter only the necessary fields
	var devices []models.DeviceResponse
	for i := range refreshTokens {
		devices = append(devices, models.DeviceResponse{
			ID:         refreshTokens[i].ID,
			LastActive: refreshTokens[i].CreateTime,
			UserAgent:  refreshTokens[i].UserAgent,
			IPAddress:  refreshTokens[i].IPAddress,
		})
	}

	currentRefreshToken, err := database.GetRefreshTokenByAccessToken(ctx.Request.Context(), tokenString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting current device"})
		return
	}

	var deviceList models.DeviceListResponse
	deviceList.CurrentDeviceID = currentRefreshToken.ID
	deviceList.Devices = devices

	ctx.JSON(http.StatusOK, deviceList)
}
