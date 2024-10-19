package config

import "rest/models"

func MigrateDB() {
	DB.AutoMigrate(&models.Posts{}, &models.Comments{})

	// Add table suffix when creating tables
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Posts{}, &models.Comments{})
}
