package main

import (
	"AltTube-Go/auth"
	"AltTube-Go/database"
	docs "AltTube-Go/docs"
	"AltTube-Go/handlers"
	"AltTube-Go/handlers/piped"
	"AltTube-Go/handlers/piped/opensearch"
	"AltTube-Go/handlers/pipedproxy"
	"AltTube-Go/handlers/user_handlers"
	"AltTube-Go/handlers/user_handlers/devices"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
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
	r.GET("/piped/opensearch/suggestions", opensearch.Suggestions)
	r.GET("/piped/search", piped.Search)
	r.GET("/piped/streams/:videoID", piped.Streams)
	r.GET("/pipedproxy/*action", pipedproxy.PipedProxy)
	user := r.Group("/user")
	{
		user.POST("/login", user_handlers.Login)
		user.POST("/signup", user_handlers.Signup)
		user.PATCH("/email", auth.Middleware(), user_handlers.EditEmail)
		user.DELETE("/", auth.Middleware(), user_handlers.DeleteUser)
		user.POST("/logout", auth.Middleware(), user_handlers.LogoutUser)
		user.POST("/refresh_token", user_handlers.RefreshToken)
		user.GET("/devices", auth.Middleware(), devices.GetDevices)
		user.DELETE("/devices", auth.Middleware(), devices.DeleteDevices)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("API failed to start: %v", err)
	}
}
