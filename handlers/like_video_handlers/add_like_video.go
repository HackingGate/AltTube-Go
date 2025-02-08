package like_video_handlers

import (
	"AltTube-Go/database"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddLikeVideo godoc
// @Summary Add like to video
// @Description Add like to video
// @Tags like
// @Accept  json
// @Produce  json
// @Param videoID path string true "Video ID"
// @Success 200 {string} JSON "{"message": "Video liked successfully"}"
// @Security AccessToken
// @Router /like/{videoID} [post]
func AddLikeVideo(ctx *gin.Context) {
	videoID := ctx.Param("videoID")

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

	if isLiked {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Video already liked"})
		return
	}

	err = database.AddLikeVideo(ctx.Request.Context(), user.ID, video.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding like video"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Video liked successfully"})
}
