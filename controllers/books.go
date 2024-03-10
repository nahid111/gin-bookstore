package controllers

import (
	"net/http"

	"gin-bookstore/database"
	"github.com/gin-gonic/gin"
)

// GET /books
func FindBooks(c *gin.Context) {
	var books []database.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books, "count": len(books)})
}

// GET /books/:id
func FindBook(c *gin.Context) {
	var book database.Book
	err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
func CreateBook(c *gin.Context) {
	// Validate input
	var input database.CreateBookSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	user := GetCurrentUser(c)
	book := database.Book{Title: input.Title, UserID: user.ID}
	database.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
func UpdateBook(c *gin.Context) {
	var book database.Book
	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input database.UpdateBookSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
func DeleteBook(c *gin.Context) {
	var book database.Book
	err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	database.DB.Delete(&book)
	// recordId := c.Param("id")
	// err := database.DB.Delete(&database.Book{}, recordId)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"data": true})
}
