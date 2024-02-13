package handlers

import "github.com/gin-gonic/gin"

func Resource(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "resource data",
	})
}
