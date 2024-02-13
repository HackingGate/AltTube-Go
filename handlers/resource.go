package handlers

import (
	"AltTube-Go/auth"
	"AltTube-Go/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func Resource(ctx *gin.Context) {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if bearerToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	reqToken := strings.Split(bearerToken, " ")[1]
	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.GetJwtKey(), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": "resource data",
	})
}
