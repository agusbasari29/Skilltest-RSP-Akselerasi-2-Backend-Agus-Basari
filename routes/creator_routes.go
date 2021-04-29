package routes

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/handler"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/middleware"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

type CreatorRoutes struct{}

func (r CreatorRoutes) Route() []helper.Route {
	db := database.SetupDatabaseConnection()
	userRepo := repository.NewUserRepository(db)
	userServices := services.NewUserServices(userRepo)
	jwtServices := services.NewJWTService()
	creatorHandler := handler.NewCreatorHandler(userServices, jwtServices)
	cache := middleware.CacheCheck()

	return []helper.Route{
		{
			Method:  "GET",
			Path:    "/creators",
			Handler: []gin.HandlerFunc{cache, creatorHandler.GetAllCreator},
		},
		{
			Method:  "POST",
			Path:    "/creators",
			Handler: []gin.HandlerFunc{creatorHandler.CreateCreator},
		},
		{
			Method:  "PUT",
			Path:    "/creators",
			Handler: []gin.HandlerFunc{creatorHandler.UpdateCreator},
		},
		{
			Method:  "DELETE",
			Path:    "/creators",
			Handler: []gin.HandlerFunc{creatorHandler.DeleteCreator},
		},
	}
}
