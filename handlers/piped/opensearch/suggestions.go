package opensearch

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func Suggestions(ctx *gin.Context) {
	// Retrieve the backend URL from an environment variable
	backendURL := os.Getenv("PIPED_BACKEND_URL")
	if backendURL == "" {
		// Handle the case where the environment variable is not set
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Backend URL is not set"})
		return
	}

	// Retrieve the query parameter from the request
	query := ctx.Query("query")
	if query == "" {
		// Handle the case where the query parameter is missing
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	// Build the request URL with the query parameter
	requestURL := backendURL + "/opensearch/suggestions?query=" + query

	// Make the HTTP GET request to the backend
	resp, err := http.Get(requestURL)
	if err != nil {
		// Handle any error that occurred during the HTTP request
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to request backend"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle any error that occurred while reading the response body
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from backend"})
		return
	}

	// Return the response body as is
	ctx.Data(resp.StatusCode, "application/json", body)
}
