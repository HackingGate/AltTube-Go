package like_video_handlers

import (
	"AltTube-Go/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLikeVideo godoc
// @Summary Get like status of video
// @Description Get like status of video
// @Tags like
// @Accept  json
// @Produce  json
// @Param videoID path string true "Video ID"
// @Success 200 {string} JSON "{"is_liked": "true"}"
// @Security AccessToken
// @Router /like/{videoID} [get]
func GetLikeVideo(ctx *gin.Context) {
	videoID := ctx.Param("videoID")

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

	user, err := database.GetUserByID(ctx.Request.Context(), authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}

	video, err := database.GetVideoByV(ctx.Request.Context(), videoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting video"})
		return
	}

	isLiked, err := database.ReadIsLikedVideo(ctx.Request.Context(), user.ID, video.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading if video is liked"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"is_liked": isLiked})
}
