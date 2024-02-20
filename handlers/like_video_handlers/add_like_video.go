package like_video_handlers

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddLike godoc
// @Summary Add like to video
// @Description Add like to video
// @Tags like
// @Accept  json
// @Produce  json
// @Param likeVideoRequest body models.LikeVideoRequest true "Like Video Request"
// @Success 200 {string} JSON "{"message": "Video liked successfully"}"
// @Security AccessToken
// @Router /like/video [post]
func AddLike(ctx *gin.Context) {
	var likeVideoRequest models.LikeVideoRequest
	if err := ctx.ShouldBindJSON(&likeVideoRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	video, err := database.GetVideoByV(likeVideoRequest.VideoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting video"})
		return
	}

	isLiked, err := database.ReadIsLikedVideo(user, video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading if video is liked"})
		return
	}

	if isLiked {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Video already liked"})
		return
	}

	err = database.AddLikeVideo(user, video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding like video"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Video liked successfully"})
}
