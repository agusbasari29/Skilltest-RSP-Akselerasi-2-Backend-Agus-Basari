package routes

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/middleware"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

func DefineAuthApiRoutes(e *gin.Engine) {
	handlers := []helper.Handler{
		AuthRoutes{},
	}
	var routes []helper.Route
	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}
	api := e.Group("/auth")
	for _, route := range routes {
		api.POST(route.Path, route.Handler...)
	}
}

func DefineSecureApiRoutes(e *gin.Engine) {
	var jwtService services.JWTServices = services.NewJWTService()
	handlers := []helper.Handler{
		UserRoute{},
		TransactionRoutes{},
		EventRoutes{},
		CreatorRoutes{},
		ReportRoutes{},
	}

	var routes []helper.Route

	for _, handler := range handlers {
		routes = append(routes, handler.Route()...)
	}

	api := e.Group("/api", middleware.AuthorizeJWT(jwtService))

	for _, route := range routes {
		switch route.Method {
		case "POST":
			{
				api.POST(route.Path, route.Handler...)
			}
		case "GET":
			{
				api.GET(route.Path, route.Handler...)
			}
		case "PUT":
			{
				api.PUT(route.Path, route.Handler...)
			}
		case "PATCH":
			{
				api.PATCH(route.Path, route.Handler...)
			}
		case "DELETE":
			{
				api.DELETE(route.Path, route.Handler...)
			}
		}
	}
}
