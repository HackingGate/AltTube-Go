package like_video_handlers

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLikedVideos godoc
// @Summary Get liked videos
// @Description Get liked videos
// @Tags like
// @Accept  json
// @Produce  json
// @Success 200 {array} models.LikeVideoResponse
// @Security AccessToken
// @Router /like/ [get]
func GetLikedVideos(ctx *gin.Context) {
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

	allLikes, err := database.GetAllLikesByUserID(authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting liked videos"})
		return
	}

	// Map allLikes to videos
	var videos []*models.Video
	for _, like := range allLikes {
		video, err := database.GetVideoByV(like.VideoID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting video"})
			return
		}
		videos = append(videos, video)
	}

	// Map videos to LikeVideoResponse
	var likeVideosResponse []models.LikeVideoResponse
	for _, video := range videos {
		likeVideosResponse = append(likeVideosResponse, models.LikeVideoResponse{
			ID:           video.ID,
			Title:        video.Title,
			ThumbnailUrl: video.ThumbnailUrl,
		})
	}

	ctx.JSON(http.StatusOK, likeVideosResponse)
}
