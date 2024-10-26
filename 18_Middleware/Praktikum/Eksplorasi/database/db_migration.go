package database

import "rest/models"

func MigrateDB() {
	DB.AutoMigrate(&models.Users{}, &models.Posts{}, &models.Categories{})

	// Add table suffix when creating tables
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Users{}, &models.Posts{}, &models.Categories{})
}
