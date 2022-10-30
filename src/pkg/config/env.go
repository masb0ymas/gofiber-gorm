package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*
Declare env with default value

example: Env("APP_NAME", "any value")
*/
func Env(key string, fallback string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if value, ok := os.LookupEnv(key); ok {
		// Return the value of the variable
		return value
	}

	// Return the value of the variable
	return fallback
}
