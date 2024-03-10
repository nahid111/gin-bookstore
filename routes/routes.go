package routes

import (
	"github.com/gin-gonic/gin"

	"gin-bookstore/controllers"
	"gin-bookstore/middlewares"
)

func SetupRoutes(r *gin.Engine) {

	// Authentication routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("", controllers.LoginHandler)
	}

	// Books routes
	booksGroup := r.Group("/books")
	{
		booksGroup.GET("", controllers.FindBooks)
		booksGroup.GET("/:id", controllers.FindBook)
		booksGroup.POST("", middlewares.AuthMiddleware(), controllers.CreateBook)
		booksGroup.PATCH("/:id", middlewares.AuthMiddleware(), controllers.UpdateBook)
		booksGroup.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteBook)
	}

	// Users routes
	usersGroup := r.Group("/users")
	{
		usersGroup.GET("", controllers.FindUsers)
		usersGroup.GET("/:id", controllers.FindUser)
		usersGroup.POST("", controllers.CreateUser)
		usersGroup.PATCH("/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
		usersGroup.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)
		usersGroup.GET("/current", middlewares.AuthMiddleware(), controllers.FindCurrentUser)
	}
}
