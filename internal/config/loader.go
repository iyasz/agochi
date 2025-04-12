package config

import (
	"os"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/utils"
	"github.com/joho/godotenv"
)

func LoadDB() *Config {

	err := godotenv.Load()

	if err != nil {
		utils.Log.Error("Error loading .env file", "Error", err)
	}
	

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			Name:     os.Getenv("DATABASE_NAME"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Timezone: os.Getenv("DATABASE_TIMEZONE"),
		},
		Jwt: Jwt{
			RefreshKey: os.Getenv("JWT_REFRESH_KEY"),
			AccessKey:  os.Getenv("JWT_ACCESS_KEY"),
		},
	}
}
