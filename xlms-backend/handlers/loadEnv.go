package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AccessTokenSecret  string
	RefreshTokenSecret string
	DB_URL             string
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}

	AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")
	DB_URL = os.Getenv("DB_URL")

	if AccessTokenSecret == "" {
		return fmt.Errorf("ACCESS_TOKEN_SECRET not set")
	}
	if RefreshTokenSecret == "" {
		return fmt.Errorf("REFRESH_TOKEN_SECRET not set")
	}
	if DB_URL == "" {
		return fmt.Errorf("DB_URL not set")
	}

	return nil
}
