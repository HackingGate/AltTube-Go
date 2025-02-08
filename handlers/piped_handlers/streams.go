package piped_handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/hackinggate/alttube-go/database"
	"github.com/hackinggate/alttube-go/models"
	"github.com/hackinggate/alttube-go/utils"

	"github.com/gin-gonic/gin"
)

// Streams godoc
// @Summary Get video streams
// @Description Get video streams by video ID
// @Tags piped
// @Accept  json
// @Produce  json
// @Param videoID path string true "Video ID"
// @Success 200 {string} JSON "Video streams"
// @Router /piped/streams/{videoID} [get]
func Streams(ctx *gin.Context) {
	// Retrieve the backend URL from an environment variable
	backendURL := os.Getenv("PIPED_BACKEND_URL")
	if backendURL == "" {
		// Handle the case where the environment variable is not set
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Backend URL is not set"})
		return
	}

	// Retrieve the query parameter from the request
	videoID := ctx.Param("videoID")
	if videoID == "" {
		// Handle the case where the stream ID is missing
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Video ID is required"})
		return
	}

	// Build the request URL with the query parameter
	requestURL := backendURL + "/streams/" + videoID

	// Make the HTTP GET request to the backend
	resp, err := http.Get(requestURL)
	if err != nil {
		// Handle any error that occurred during the HTTP request
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to request backend"})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close response body"})
			return
		}
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle any error that occurred while reading the response body
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from backend"})
		return
	}

	// Call the RewriteURLsInJSON utility
	modifiedBody, err := utils.RewriteURLsInJSONStringBased(body, os.Getenv("PIPED_PROXY_URL_FOR_REWRITE"), "/pipedproxy")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to modify URLs in response"})
		return
	}

	// Add Video to database if it doesn't exist
	if resp.StatusCode == 200 {
		var video models.Video
		// Decode JSON and store in video
		err := json.Unmarshal(modifiedBody, &video)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal response from backend"})
			return
		}

		video.ID = videoID

		// Check if video already exists in the database
		existingVideo := database.VideoExists(video.ID)

		if !existingVideo {
			// Save the new video to the database
			err = database.AddVideo(video)
			if err != nil {
				// Handle potential database error
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video to database"})
				return
			}
		}
	}

	// Return the response body as is
	ctx.Data(resp.StatusCode, "application/json", modifiedBody)
}
