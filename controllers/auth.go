package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-bookstore/database"
	"gin-bookstore/utils"
)

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
