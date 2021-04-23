package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConnection struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
}

func DbConfig() DbConnection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	dbConfig := DbConnection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	return dbConfig
}
