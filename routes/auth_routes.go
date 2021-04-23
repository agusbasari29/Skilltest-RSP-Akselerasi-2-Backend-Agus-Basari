package routes

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/handler"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct{}

func (r AuthRoutes) Route() []helper.Route {
	db := database.SetupDatabaseConnection()
	useRepo := repository.NewUserRepository(db)
	authServices := services.NewAuthService(useRepo)
	jwtService := services.NewJWTService()
	authHandler := handler.NewAuthHandler(authServices, jwtService)

	return []helper.Route{
		{
			Method:  "POST",
			Path:    "/register",
			Handler: []gin.HandlerFunc{authHandler.Register},
		},
		{
			Method:  "POST",
			Path:    "/login",
			Handler: []gin.HandlerFunc{authHandler.Login},
		},
		{
			Method:  "POST",
			Path:    "/forget_password",
			Handler: []gin.HandlerFunc{authHandler.ForgetPassword},
		},
	}
}
