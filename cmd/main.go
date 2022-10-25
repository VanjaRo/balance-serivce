package main

import (
	"os"

	"github.com/VanjaRo/balance-serivce/pkg/api"
	"github.com/joho/godotenv"
)

func main() {
	// init env variables
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	api.Start(&api.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),
	})
}
