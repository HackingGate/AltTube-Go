package piped

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

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

	// Return the response body as is
	ctx.Data(resp.StatusCode, "application/json", body)
}
