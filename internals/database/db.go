package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// we use a GORM for connecting golang with postgresql
var DB *gorm.DB

// Connect initializes the database connection using GORM and PostgreSQL
func Connect() {
	if os.Getenv("DATABASE_USER") == "" || os.Getenv("DATABASE_PASSWORD") == "" || os.Getenv("DATABASE_NAME") == "" {
		fmt.Println("Database Environment variables is Empty")
		return
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to databse \n", err)
	}
	DB = database
	fmt.Println("Database Connected Successfully!")
}
