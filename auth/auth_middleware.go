package auth

import (
	"AltTube-Go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
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

		// If token is not in tokens, it is logged out
		if !contains(tokens, tokenString) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// Token is valid, set email in context, so it can be used in the handler
		ctx.Set("uuid", claims.UUID)
		ctx.Next()
	}
}

// contains checks if the tokens slice contains a specific token.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
