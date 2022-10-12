package database

import (
	"fmt"
	"log"

	"blitzomni.com/m/models"
	"blitzomni.com/m/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *utils.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println("🚀 Connected successfully to the database")

	return db
}

func init() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	DB = ConnectDB(&config)

	DB.AutoMigrate(&models.User{})
	fmt.Println("👍 Migration complete")
}
