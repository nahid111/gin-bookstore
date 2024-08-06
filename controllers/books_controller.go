package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-bookstore/crud"
	"gin-bookstore/database"
)

// @Summary List books
// @Description Get a list of books with pagination
// @Tags books
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 10)"
// @Success 200 {object} PaginatedBooksResponse
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

// @Summary Get a book by ID
// @Description Retrieve a book's details by its ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} database.Book
// @Failure 400 {object} ErrorResponse
// @Router /books/{id} [get]
func FindBook(c *gin.Context) {
	book, err := crud.ReadBookBy("id", c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// @Summary Create a new book
// @Description Add a new book to the collection
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param book body database.CreateBookSchema true "Book data"
// @Success 200 {object} database.Book
// @Failure 400 {object} ErrorResponse
// @Router /books [post]
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

// @Summary Update a book by ID
// @Description Modify the details of a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Param book body database.UpdateBookSchema true "Updated book data"
// @Success 200 {object} database.Book
// @Failure 400 {object} ErrorResponse
// @Router /books/{id} [put]
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

// @Summary Delete a book by ID
// @Description Remove a book from the collection by its ID
// @Tags books
// @Produce json
// @Security BearerAuth
// @Param id path string true "Book ID"
// @Success 200 {object} database.Book
// @Failure 400 {object} ErrorResponse
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	book, err := crud.DeleteBookByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PaginatedBooksResponse represents the structure of the response for paginated books
// @Description Response structure for paginated list of books
// @Param data query []database.Book true "List of books"
// @Param total_count query int true "Total number of books"
// @Param page query int true "Current page number"
// @Param page_size query int true "Number of items per page"
// @Param total_pages query int true "Total number of pages"
// @Success 200 {object} PaginatedBooksResponse
type PaginatedBooksResponse struct {
	Data        []database.Book `json:"data"`
	TotalCount  int             `json:"total_count"`
	Page        int             `json:"page"`
	PageSize    int             `json:"page_size"`
	TotalPages  int             `json:"total_pages"`
}
