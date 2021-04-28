package routes

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/handler"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

type ReportRoutes struct{}

func (r ReportRoutes) Route() []helper.Route {
	db := database.SetupDatabaseConnection()
	reportRepo := repository.NewReportRepository(db)
	reportServices := services.NewReportService(reportRepo)
	jwt := services.NewJWTService()
	reportHandler := handler.NewReportHandler(reportServices, jwt)

	return []helper.Route{
		{
			Path:    "/detail_report",
			Method:  "POST",
			Handler: []gin.HandlerFunc{reportHandler.DetailReportByEvent},
		}, {
			Path:    "/summary_report",
			Method:  "GET",
			Handler: []gin.HandlerFunc{reportHandler.AllSummaryReport},
		}, {
			Path:    "/creator_summary_report",
			Method:  "GET",
			Handler: []gin.HandlerFunc{reportHandler.AllSummaryReportByCreator},
		},
	}
}
