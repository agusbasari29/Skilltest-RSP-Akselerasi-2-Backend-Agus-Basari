package seeders

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB = database.SetupDatabaseConnection()
	eventRepo          = repository.NewEventRepository(db)
	userRepo           = repository.NewUserRepository(db)
	trxRepo            = repository.NewTransactionRepository(db)
)

func Seeders() {
	UsersSeedersUp(35)
	EventsSeedersUp(100)
	TransactionsSeedersUp(2000)
}
