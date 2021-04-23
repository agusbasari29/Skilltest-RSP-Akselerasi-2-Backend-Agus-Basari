package main

import (
	"os"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDatabaseConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)
	g := gin.Default()
	g.GET("/init", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "project sucessfully initiated..."})
	})
	routes.DefineAuthApiRoutes(g)
	routes.DefineSecureApiRoutes(g)
	g.Run(os.Getenv("SERVER_PORT"))
}
