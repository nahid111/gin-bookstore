package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-bookstore/database"
	"gin-bookstore/utils"
)

// LoginHandler handles user login requests.
//
// @Summary      User Login
// @Description  Logs in a user with their email and password, returns an authentication token on success.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        loginBody body database.CreateUserSchema true "Login credentials"
// @Success      200  {object}  map[string]string "Successful login with token"
// @Header       200  {string}  Authorization "Bearer <token>"  // Specify that a JWT token is returned in the Authorization header
// @Failure      400  {object}  map[string]string "Invalid Credentials"
// @Router       /auth [post]
func LoginHandler(c *gin.Context) {
	// Validate input
	var input database.CreateUserSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check User
	var user database.User
	result := database.DB.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials!"})
		return
	}

	// Check Password
	if !utils.ComparePasswords(user.Password, input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials!"})
		return
	}

	token := utils.GenerateAuthToken(input.Email, user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
