package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// we use a GORM for connecting golang with postgresql
var DB *gorm.DB

// Connect initializes the database connection using GORM and PostgreSQL
func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env files")
		return
	}
	if os.Getenv("DATABASE_USER") == "" || os.Getenv("DATABASE_PASSWORD") == "" || os.Getenv("DATABASE_NAME") == "" {
		fmt.Println("Database Environment is Empty")
		return
	}
	fmt.Println(os.Getenv("DATABASE_USER"))
	ds := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))
	database, err := gorm.Open(postgres.Open(ds), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to databse \n", err)
	}
	DB = database
	fmt.Println("Database Connected Successfully!")

}
