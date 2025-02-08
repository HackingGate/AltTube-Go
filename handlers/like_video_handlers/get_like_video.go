package like_video_handlers

import (
	"net/http"

	"github.com/hackinggate/alttube-go/database"

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

	user, err := database.GetUserByID(authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}

	video, err := database.GetVideoByV(videoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting video"})
		return
	}

	isLiked, err := database.ReadIsLikedVideo(user, video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading if video is liked"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"is_liked": isLiked})
}
