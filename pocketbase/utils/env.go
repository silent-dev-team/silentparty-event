package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found - ignore on docker deployment")
	}
	return os.Getenv(key)
}
