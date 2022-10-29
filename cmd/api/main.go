package main

import (
	"os"

	"github.com/VanjaRo/balance-serivce/pkg/api"
	"github.com/joho/godotenv"
)

const (
	servicesProfile = "SERVICES_PROFILE"
	dockerProfile   = "docker"
)

func main() {
	// init env variables
	profile := os.Getenv(servicesProfile)
	if profile != dockerProfile {
		if err := godotenv.Load("local.env"); err != nil {
			panic("Error loading .env file")
		}
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
