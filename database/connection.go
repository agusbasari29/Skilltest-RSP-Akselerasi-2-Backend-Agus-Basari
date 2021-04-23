package database

import (
	"fmt"
	"log"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	dbConfig := config.DbConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DbName,
		dbConfig.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed connecting to database!")
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Failed to close connection database.")
	}
	conn.Close()
}
