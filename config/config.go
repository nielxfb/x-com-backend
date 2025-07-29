package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")

	if DBHost == "" || DBPort == "" || DBUser == "" || DBPass == "" || DBName == "" {
		log.Fatal("Database configuration is not set properly. Please check your environment variables.")
	}
}
