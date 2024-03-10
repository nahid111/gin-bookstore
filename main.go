package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"gin-bookstore/database"
	"gin-bookstore/routes"
)

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

	// Run the server
	r.Run()
}
