package main

import (
	"AltTube-Go/database"
	"AltTube-Go/handlers"
	"AltTube-Go/handlers/piped"
	"AltTube-Go/handlers/piped/opensearch"
	"AltTube-Go/handlers/pipedproxy"
	"AltTube-Go/handlers/resource"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var r *gin.Engine

func main() {
	database.Init()
	startApi()
}

func startApi() {
	r = gin.Default()

	r.GET("/ping", handlers.Ping)
	r.GET("/piped/opensearch/suggestions", opensearch.Suggestions)
	r.GET("/piped/search", piped.Search)
	r.GET("/streams/:videoID", piped.Streams)
	r.GET("/pipedproxy/*action", pipedproxy.PipedProxy)
	r.GET("/resource", resource.Resource)
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("API failed to start: %v", err)
	}
}
