package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-bookstore/database"
	"gin-bookstore/utils"
)

func GetCurrentUser(c *gin.Context) *database.User {
	var user database.User
	email := c.MustGet("currentUserEmail")
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println("Error fetching User", err)
		return nil
	}
	return &user
}

// GET /users
func FindUsers(c *gin.Context) {
	var users []database.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users, "count": len(users)})
}

// GET /users/current
func FindCurrentUser(c *gin.Context) {
	user := GetCurrentUser(c)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users/:id
func FindUser(c *gin.Context) {
	var user database.User
	err := database.DB.Where("id = ?", c.Param("id")).Preload("Books").First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /users
func CreateUser(c *gin.Context) {
	// Validate input
	var input database.CreateUserSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	var user database.User
	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Exists"})
		return
	}

	// Create User
	hashedPass := utils.HashPassword(input.Password)
	if hashedPass == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed!"})
		return
	}

	user = database.User{Email: input.Email, Password: *hashedPass}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/:id
func UpdateUser(c *gin.Context) {
	var user database.User
	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input database.UpdateUserSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	var user database.User
	err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
