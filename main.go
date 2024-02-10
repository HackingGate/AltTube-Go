package main

import (
	"AltTube-Go/database"
	"AltTube-Go/handlers"
	"AltTube-Go/handlers/piped"
	"AltTube-Go/handlers/piped/opensearch"
	"AltTube-Go/handlers/pipedproxy"
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
	r.GET("/pipedproxy/*action", pipedproxy.PipedProxy)
	if err := r.Run(); err != nil {
		log.Fatalf("API failed to start: %v", err)
	}
}
