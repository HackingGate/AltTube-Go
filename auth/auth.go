package auth

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

var tokens []string

func GetTokens() []string {
	return tokens
}

func AddToken(token string) {
	tokens = append(tokens, token)
}

func RandomHex(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	})
}
