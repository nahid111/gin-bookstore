package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-bookstore/crud"
	"gin-bookstore/database"
)

// GET /users
func FindUsers(c *gin.Context) {
	users, err := crud.ReadUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users, "count": len(users)})
}

// GET /users/current
func FindCurrentUser(c *gin.Context) {
	email := c.MustGet("currentUserEmail")
	user, err := crud.ReadUserBy("email", email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users/:id
func FindUser(c *gin.Context) {
	user, err := crud.ReadUserBy("id", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	// Create User
	user, err := crud.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/:id
func UpdateUser(c *gin.Context) {
	// Validate input
	var input database.UpdateUserSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update User
	user, err := crud.UpdateUserById(c.Param("id"), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	user, err := crud.DeleteUserByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
