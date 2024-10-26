package database

import (
	"fmt"
	"rest/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetConfig("DB_USERNAME"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_PORT"),
		utils.GetConfig("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
