package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Return the value of the variable
	return os.Getenv(key)
}
