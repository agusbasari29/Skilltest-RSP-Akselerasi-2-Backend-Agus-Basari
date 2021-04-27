package main

import (
	"net/http"
	"os"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database/seeders"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDatabaseConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)
	db.AutoMigrate(&entity.Users{}, &entity.Event{}, &entity.Transaction{})
	g := gin.Default()
	g.GET("/seeder", func(ctx *gin.Context) {
		seeders.Seeders()
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully seed database"})
	})
	routes.DefineAuthApiRoutes(g)
	routes.DefineSecureApiRoutes(g)
	g.Run(os.Getenv("SERVER_PORT"))
}
