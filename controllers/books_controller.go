package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-bookstore/crud"
	"gin-bookstore/database"
)

// GET /books
func FindBooks(c *gin.Context) {
	books, err := crud.ReadBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books, "count": len(books)})
}

// GET /books/:id
func FindBook(c *gin.Context) {
	book, err := crud.ReadBookBy("id", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	// Get current user email
	email := c.MustGet("currentUserEmail")

	// Create book
	book, err := crud.CreateBook(input, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
func UpdateBook(c *gin.Context) {
	// Validate input
	var input database.UpdateBookSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update Book
	book, err := crud.UpdateBookById(c.Param("id"), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
func DeleteBook(c *gin.Context) {
	book, err := crud.DeleteBookByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}
