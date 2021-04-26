package routes

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/handler"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

type TransactionRoutes struct{}

func (r TransactionRoutes) Route() []helper.Route {
	db := database.SetupDatabaseConnection()
	trxRepo := repository.NewTransactionRepository(db)
	trxServices := services.NewTransactionServices(trxRepo)
	jwtServices := services.NewJWTService()
	participantHandler := handler.NewParticipantHandler(trxServices, jwtServices)

	return []helper.Route{
		{
			Path:    "/payment_confirmation",
			Method:  "POST",
			Handler: []gin.HandlerFunc{participantHandler.GetPasticipantPendingTransaction},
		},
		{
			Path:    "/pending_payment",
			Method:  "GET",
			Handler: []gin.HandlerFunc{participantHandler.GetAllPendingTransaction},
		},
		{
			Path:    "/payment_status",
			Method:  "POST",
			Handler: []gin.HandlerFunc{participantHandler.ChangeStatusPaymentParticipant},
		},
	}
}
