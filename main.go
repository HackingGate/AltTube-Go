package main

import (
	"AltTube-Go/database"
	"AltTube-Go/handlers"
	"AltTube-Go/handlers/piped"
	"AltTube-Go/handlers/piped/opensearch"
	"github.com/gin-gonic/gin"
	"log"
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
	if err := r.Run(); err != nil {
		log.Fatalf("API failed to start: %v", err)
	}
}
