package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-bookstore/crud"
	"gin-bookstore/database"
)

// @Summary list books
// @Description get list of books
// @Tags books
// @Produce json
// @Success 200 {array} database.Book
// @Router /books [get]
func FindBooks(c *gin.Context) {
	// Get page and pageSize from query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// Ensure page and pageSize are valid
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	paginatedBooks, err := crud.ReadBooks(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        paginatedBooks.Books,
		"total_count": paginatedBooks.TotalCount,
		"page":        paginatedBooks.Page,
		"page_size":   paginatedBooks.PageSize,
		"total_pages": paginatedBooks.TotalPages,
	})
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
