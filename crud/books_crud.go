package crud

import (
	"fmt"

	"gorm.io/gorm/clause"

	"gin-bookstore/database"
	"gin-bookstore/logger"
)

func ReadBooks() ([]database.Book, error) {
	var Books []database.Book
	result := database.DB.Preload(clause.Associations).Find(&Books)
	if result.Error != nil {
		logger.Error.Println(result.Error)
		return nil, result.Error
	}
	return Books, nil
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
