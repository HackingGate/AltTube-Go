package user_handlers

import (
	"net/http"

	"github.com/hackinggate/alttube-go/database"
	"github.com/hackinggate/alttube-go/models"
	"github.com/hackinggate/alttube-go/utils"

	"github.com/gin-gonic/gin"
)

// Signup godoc
// @Summary Signup
// @Description Signup
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.SignupRequest true "User"
// @Success 200 {string} JSON "{"message": "Registration successful"}"
// @Router /user/signup [post]
func Signup(ctx *gin.Context) {
	var signupRequest models.SignupRequest
	if err := ctx.ShouldBindJSON(&signupRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := utils.HashPassword(signupRequest.Password) // Hash the password
	signupRequest.Password = hashedPassword

	_, err := database.AddUser(ctx.Request.Context(), signupRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
