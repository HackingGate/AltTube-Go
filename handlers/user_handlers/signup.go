package user_handlers

import (
	"AltTube-Go/database"
	"AltTube-Go/models"
	"AltTube-Go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Signup godoc
// @Summary Signup
// @Description Signup
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.Signup true "User"
// @Success 200 {string} JSON "{"message": "Registration successful"}"
// @Router /user/signup [post]
func Signup(c *gin.Context) {
	var signup models.Signup
	if err := c.ShouldBindJSON(&signup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user.Email = signup.Email
	user.Password = signup.Password
	user.UserID = utils.GenerateUUID()

	hashedPassword, _ := utils.HashPassword(user.Password) // Hash the password
	user.Password = hashedPassword

	err := database.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
