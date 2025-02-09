package auth

import (
	"net/http"
	"strings"

	"github.com/hackinggate/alttube-go/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		// Assuming token is provided as 'Bearer <token>'
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &tokenClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "JWT error"})
			ctx.Abort()
			return
		}

		exists, err := database.ValidateAccessToken(ctx.Request.Context(), tokenString)
		if !exists {
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
