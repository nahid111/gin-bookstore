package crud

import (
	"fmt"

	"gorm.io/gorm/clause"

	"gin-bookstore/database"
	"gin-bookstore/logger"
)

// PaginatedBooks represents the paginated response
type PaginatedBooks struct {
	Books      []database.Book `json:"books"`
	TotalCount int64           `json:"total_count"`
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalPages int             `json:"total_pages"`
}

func ReadBooks(page, pageSize int) (PaginatedBooks, error) {
	var books []database.Book
	var totalCount int64

	// Count total number of books
	if err := database.DB.Model(&database.Book{}).Count(&totalCount).Error; err != nil {
		logger.Error.Println(err)
		return PaginatedBooks{}, err
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Fetch books with pagination
	result := database.DB.Preload(clause.Associations).
		Offset(offset).
		Limit(pageSize).
		Find(&books)

	if result.Error != nil {
		logger.Error.Println(result.Error)
		return PaginatedBooks{}, result.Error
	}

	// Calculate total pages
	totalPages := int(totalCount) / pageSize
	if int(totalCount)%pageSize != 0 {
		totalPages++
	}

	paginatedResponse := PaginatedBooks{
		Books:      books,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}

	return paginatedResponse, nil
}

func ReadBookBy(field string, value any) (database.Book, error) {
	var Book database.Book
	q := fmt.Sprintf("%s = ?", field)
	result := database.DB.Preload("User").Where(q, value).First(&Book)
	if result.Error != nil {
		logger.Error.Println(result.Error, value)
		return database.Book{}, result.Error
	}
	return Book, nil
}

func CreateBook(input database.CreateBookSchema, email any) (database.Book, error) {
	// fetch user from email
	user, _ := ReadUserBy("email", email)
	// Create Book
	book := database.Book{Title: input.Title, UserID: user.ID}
	database.DB.Create(&book)
	return book, nil
}

func UpdateBookById(id string, input database.UpdateBookSchema) (database.Book, error) {
	// fetching Book
	var Book database.Book
	result := database.DB.Where("id = ?", id).First(&Book)
	if result.Error != nil {
		logger.Error.Println(result.Error)
		return database.Book{}, result.Error
	}
	// Update Book
	database.DB.Model(&Book).Updates(input)
	return Book, nil
}

func DeleteBookByID(id string) (database.Book, error) {
	var Book database.Book
	result := database.DB.Where("id = ?", id).First(&Book)
	if result.Error != nil {
		logger.Error.Println(result.Error, id)
		return database.Book{}, result.Error
	}
	database.DB.Delete(&Book)
	return Book, nil
}
