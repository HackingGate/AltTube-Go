package auth

import (
	"AltTube-Go/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func GenerateJWT() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &model.Claims{
		Username: "username",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	})
}
