package user_handlers

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"AltTube-Go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := utils.HashPassword(user.Password) // Hash the password
	user.Password = hashedPassword

	err := database.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
