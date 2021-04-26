package routes

import (
	"os"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/cache"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/handler"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

type EventRoutes struct{}

func (r EventRoutes) Route() []helper.Route {
	db := database.SetupDatabaseConnection()
	db.AutoMigrate(&entity.Event{}, &entity.Transaction{})
	eventRepo := repository.NewEventRepository(db)
	trxRepo := repository.NewTransactionRepository(db)
	eventServices := services.NewEventServices(eventRepo)
	jwtServices := services.NewJWTService()
	var cacheService cache.EventCache = cache.NewRedisCache(os.Getenv("REDIS_ADDR_PORT"), 0, 10)
	trxServices := services.NewTransactionServices(trxRepo)
	eventHandler := handler.NewEventHandler(eventServices, jwtServices, trxServices, cacheService)

	return []helper.Route{
		{
			Path:    "/events",
			Method:  "POST",
			Handler: []gin.HandlerFunc{eventHandler.CreateEvent},
		}, {
			Path:    "/events",
			Method:  "GET",
			Handler: []gin.HandlerFunc{eventHandler.GetAllEvent},
		}, {
			Path:    "/events",
			Method:  "PUT",
			Handler: []gin.HandlerFunc{eventHandler.UpdateEvent},
		}, {
			Path:    "/events",
			Method:  "DELETE",
			Handler: []gin.HandlerFunc{eventHandler.DeletedEvent},
		}, {
			Path:    "/purchase",
			Method:  "post",
			Handler: []gin.HandlerFunc{eventHandler.MakeEventPurchase},
		},
	}
}
