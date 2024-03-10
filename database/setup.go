package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_NAME = os.Getenv("DB_NAME")
var DB_USER = os.Getenv("DB_USER")
var DB_PASSWORD = os.Getenv("DB_PASSWORD")
var DB_HOST = os.Getenv("DB_HOST")
var DB_PORT = os.Getenv("DB_PORT")

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
}

func Migrate() {
	InitDB()
	DB.AutoMigrate(&Book{}, &User{})
}

func DropTables() {
	InitDB()
	err := DB.Migrator().DropTable(&Book{}, &User{})
	if err != nil {
		panic(fmt.Sprintf("Failed to drop tables: %v", err))
	}
	fmt.Printf("Tables dropped successfully...\n\n")
}
