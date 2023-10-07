package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("unable to load .env file")
	}
}

func getEnv(key string, defaultVal string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultVal
	}
	return value
}

func getEnvInt(key string, defaultVal int) int {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultVal
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultVal
	}
	return intValue
}
