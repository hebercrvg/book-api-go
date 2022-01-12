package database

import (
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func BuildDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Database: os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}
