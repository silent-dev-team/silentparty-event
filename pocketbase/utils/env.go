package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found - ignore on docker deployment")
	}
	return os.Getenv(key)
}

func GetenvInt64(key string) int64 {
	val := Getenv(key)
	if val == "" {
		return 0
	}
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Printf("failed to parse %q as int64: %v", val, err)
		return 0
	}
	return intVal
}
