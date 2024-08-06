package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	_ "gin-bookstore/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-bookstore/database"
	"gin-bookstore/routes"
)

// @title gin-bookstore
// @version 1.0
// @description REST API using gin and gorm
// @host localhost:8080
// @BasePath
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Check for cmd args
	if len(os.Args) > 1 {
		for i, arg := range os.Args[1:] {
			if arg == "--clean" {
				fmt.Printf("Arg %d: %s\n", i+1, arg)
				database.DropTables()
			}
		}
	}

	database.Migrate()

	r := gin.Default()

	// Set up routes
	routes.SetupRoutes(r)

	// docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Run the server
	r.Run()
}
