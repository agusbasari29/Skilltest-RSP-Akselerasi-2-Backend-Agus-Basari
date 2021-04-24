package handler

import (
	"fmt"
	"net/http"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type participantHandler struct {
	trxServices services.TransactionServices
	jwtService  services.JWTServices
}

func NewParticipantHandler(trxServices services.TransactionServices, jwtService services.JWTServices) *participantHandler {
	return &participantHandler{trxServices, jwtService}
}

func (h *participantHandler) ChangeStatusPaymentParticipant(ctx *gin.Context) {
	var req request.RequestParticipantTransaction
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	admin := role == string(entity.Admin)
	if !admin {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	err = ctx.ShouldBind(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid data type", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = validate.Struct(req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid input type", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var trx request.RequestTransactionUpdate
	trx.ID = req.ID
	trx.StatusPayment = req.StatusPayment
	update, err := h.trxServices.UpdateTransaction(trx)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to confirm payment", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully create new transaction.", update)
	ctx.JSON(http.StatusOK, response)
}

func (h *participantHandler) GetAllPendingTransaction(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	admin := role == string(entity.Admin)
	if !admin {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	pending, err := h.trxServices.GetPendingTrasaction()
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to retreive pending transaction data", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully create new transaction.", pending)
	ctx.JSON(http.StatusOK, response)
}
