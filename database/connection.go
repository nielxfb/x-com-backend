package database

import (
	"fmt"
	"log"

	"github.com/nielxfb/x-com-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}

	initDb()

	return db
}