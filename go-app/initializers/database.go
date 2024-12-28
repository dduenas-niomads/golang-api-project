package initializers

import (
	"fmt"
	"go-app/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"))
	
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helpers.ErrorPanic(err)

	return db
}
