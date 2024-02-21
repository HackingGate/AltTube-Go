package piped_handlers

import (
	"AltTube-Go/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// Search godoc
// @Summary Search
// @Description Search
// @Tags piped
// @Accept  json
// @Produce  json
// @Param q query string true "Query"
// @Success 200 {string} JSON "Search results"
// @Router /piped/search [get]
func Search(ctx *gin.Context) {
	// Retrieve the backend URL from an environment variable
	backendURL := os.Getenv("PIPED_BACKEND_URL")
	if backendURL == "" {
		// Handle the case where the environment variable is not set
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Backend URL is not set"})
		return
	}

	// Retrieve the query parameter from the request
	q := ctx.Query("q")
	if q == "" {
		// Handle the case where the query parameter is missing
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	// Build the request URL with the query parameter
	requestURL := backendURL + "/search?q=" + q + "&filter=all"

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

	// Return the response body as is
	ctx.Data(resp.StatusCode, "application/json", modifiedBody)
}
