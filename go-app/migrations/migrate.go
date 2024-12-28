package migrations

import (
	"go-app/models"
	"gorm.io/gorm"
)
func MigrateModels(Db *gorm.DB) {

	Db.Table("users").AutoMigrate(&models.User{})

	Db.Table("tags").AutoMigrate(&models.Tags{})
}