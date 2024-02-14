package auth

import (
	"AltTube-Go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("my_secret_key")
var tokens []string

func GetJwtKey() []byte {
	return jwtKey
}

func AddToken(token string) {
	tokens = append(tokens, token)
}

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		// Assuming token is provided as 'Bearer <token>'
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// Token is valid, set email in context, so it can be used in the handler
		ctx.Set("email", claims.Email)
		ctx.Next()
	}
}
