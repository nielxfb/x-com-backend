package config

import "os"

var (
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
)