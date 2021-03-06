package routes

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/handler"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

type UserRoute struct{}

func (r UserRoute) Route() []helper.Route {
	db := database.SetupDatabaseConnection()
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserServices(userRepo)
	jwtService := services.NewJWTService()
	userHandler := handler.NewUserHandler(userService, jwtService)

	return []helper.Route{
		{
			Path:    "/users",
			Method:  "GET",
			Handler: []gin.HandlerFunc{userHandler.Profile},
		},
		{
			Path:    "/users",
			Method:  "POST",
			Handler: []gin.HandlerFunc{userHandler.UpdateUser},
		},
	}
}
