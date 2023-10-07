package config

import (
	"fmt"
)

type AppConfig struct {
	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string
	DBType string
}

// NewSQLDBConfig returns the sql database url
func NewSQLDBConfig() string {
	loadEnv()
	cfg := &AppConfig{
		DBHost: getEnv("SDB_HOST", "localhost"),
		DBPort: getEnvInt("SDB_PORT", 5432),
		DBUser: getEnv("SDB_USER", "postgres"),
		DBPass: getEnv("SDB_PASS", "postgres"),
		DBName: getEnv("SDB_NAME", "boilerplate"),
		DBType: getEnv("SDB_TYPE", "postgres"),
	}
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%d/%s", cfg.DBType, cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	return dbURL
}

// NewNOSQLDBConfig returns the nosql database url
func NewNOSQLDBConfig() string {
	loadEnv()
	cfg := &AppConfig{
		DBHost: getEnv("DNB_HOST", "localhost"),
		DBPort: getEnvInt("NDB_PORT", 27017),
		DBUser: getEnv("NDB_USER", "root"),
		DBPass: getEnv("NDB_PASS", "rootpassword"),
		DBName: getEnv("NDB_NAME", "boilerplate"),
		DBType: getEnv("NDB_TYPE", "mongodb"),
	}
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%d/%s", cfg.DBType, cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	return dbURL
}

func NewPortConfig() string {
	loadEnv()
	port := getEnv("PORT", ":3000")
	return ":" + port
}
