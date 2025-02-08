package main

import (
	"log"
	"os"

	"github.com/hackinggate/alttube-go/auth"
	"github.com/hackinggate/alttube-go/database"
	docs "github.com/hackinggate/alttube-go/docs"
	"github.com/hackinggate/alttube-go/handlers"
	"github.com/hackinggate/alttube-go/handlers/like_video_handlers"
	"github.com/hackinggate/alttube-go/handlers/piped_handlers"
	"github.com/hackinggate/alttube-go/handlers/piped_handlers/opensearch"
	"github.com/hackinggate/alttube-go/handlers/pipedproxy"
	"github.com/hackinggate/alttube-go/handlers/user_handlers"
	"github.com/hackinggate/alttube-go/handlers/user_handlers/devices"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r *gin.Engine

// @title AltTube API
// @version 1.0
// @description This is the API documentation for the AltTube application.

// @SecurityDefinitions.apiKey AccessToken
// @in header
// @name Authorization

// @SecurityDefinitions.apiKey RefreshToken
// @in header
// @name Authorization

func main() {
	database.Init()
	startApi()
}

func startApi() {
	r = gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/ping", handlers.Ping)
	pipedApi := r.Group("/piped")
	{
		pipedApi.GET("/opensearch/suggestions", opensearch.Suggestions)
		pipedApi.GET("/search", piped_handlers.Search)
		pipedApi.GET("/streams/:videoID", piped_handlers.Streams)
	}
	pipedProxyApi := r.Group("/pipedproxy")
	{
		pipedProxyApi.GET("/*action", pipedproxy.PipedProxy)
	}
	userApi := r.Group("/user")
	{
		userApi.POST("/login", user_handlers.Login)
		userApi.POST("/signup", user_handlers.Signup)
		userApi.PATCH("/email", auth.Middleware(), user_handlers.EditEmail)
		userApi.DELETE("/", auth.Middleware(), user_handlers.DeleteUser)
		userApi.POST("/logout", auth.Middleware(), user_handlers.LogoutUser)
		userApi.POST("/refresh_token", user_handlers.RefreshToken)
		userApi.GET("/devices", auth.Middleware(), devices.GetDevices)
		userApi.DELETE("/devices", auth.Middleware(), devices.DeleteDevices)
	}
	likeApi := r.Group("/like")
	{
		likeApi.GET("/:videoID", auth.Middleware(), like_video_handlers.GetLikeVideo)
		likeApi.POST("/:videoID", auth.Middleware(), like_video_handlers.AddLikeVideo)
		likeApi.DELETE("/:videoID", auth.Middleware(), like_video_handlers.RemoveLikeVideo)
		likeApi.GET("/", auth.Middleware(), like_video_handlers.GetLikeVideo)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("API failed to start: %v", err)
	}
}
