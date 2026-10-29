package config

import (
	"fmt"
	"golang-dasar/campaign"
	"golang-dasar/user"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed  connect to database", err)
	}

	err = database.AutoMigrate(&user.User{}, &campaign.Campaign{}, &campaign.CampaignImage{})
	if err != nil {
		log.Fatal("failed migrate to database", err)
	}

	DB = database
	log.Print("success connect to database")
}
