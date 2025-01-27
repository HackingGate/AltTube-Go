package auth

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		// Assuming token is provided as 'Bearer <token>'
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		_, err = database.ValidateAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			ctx.Abort()
			return
		}

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// Token is valid, set UserID in context, so it can be used in the handler
		ctx.Set("UserID", claims.UserID)
		ctx.Next()
	}
}
