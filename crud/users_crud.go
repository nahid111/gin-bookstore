package crud

import (
	"errors"
	"fmt"

	"gin-bookstore/database"
	"gin-bookstore/logger"
	"gin-bookstore/utils"
)

func ReadUsers() ([]database.User, error) {
	var users []database.User
	result := database.DB.Preload("Books").Find(&users)
	if result.Error != nil {
		logger.Error.Println(result.Error)
		return nil, result.Error
	}
	return users, nil
}

func ReadUserBy(field string, value any) (database.User, error) {
	var user database.User
	q := fmt.Sprintf("%s = ?", field)
	result := database.DB.Preload("Books").Where(q, value).First(&user)
	if result.Error != nil {
		logger.Error.Println(result.Error, value)
		return database.User{}, result.Error
	}
	return user, nil
}

func CreateUser(input database.CreateUserSchema) (database.User, error) {
	// Check if user exists
	var user database.User
	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error == nil {
		logger.Error.Println("record exists", input)
		return database.User{}, errors.New("record exists")
	}

	// Create User
	hashedPass := utils.HashPassword(input.Password)
	if hashedPass == nil {
		logger.Error.Println("password hashing failed", input)
		return database.User{}, errors.New("user creation failed")
	}

	user = database.User{Email: input.Email, Password: *hashedPass}
	database.DB.Create(&user)
	return user, nil
}

func UpdateUserById(id string, input database.UpdateUserSchema) (database.User, error) {
	// fetching user
	var user database.User
	result := database.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		logger.Error.Println(result.Error)
		return database.User{}, result.Error
	}

	// Update User
	database.DB.Model(&user).Updates(input)
	return user, nil
}

func DeleteUserByID(id string) (database.User, error) {
	var user database.User
	result := database.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		logger.Error.Println(result.Error, id)
		return database.User{}, result.Error
	}
	database.DB.Delete(&user)
	return user, nil
}
