package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	admin := role == string(entity.Admin)
	if admin {
		var req request.RequestParticipantTransaction
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
		//TODO Send Email Link
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully create new transaction.", update)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

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
	if admin {
		pending, err := h.trxServices.GetPendingTrasaction()
		if err != nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to retreive pending transaction data", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully create new transaction.", pending)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (h *participantHandler) GetPasticipantPendingTransaction(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	id, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		panic(err)
	}
	participant := role == string(entity.Participant)
	if participant {
		var req request.RequestParticipantTransaction
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
		req.ParticipantId = uint(id)
		pending, err := h.trxServices.GetParticipantPendingTrasaction(req)
		if err != nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to retreive pending transaction data", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully create new transaction.", pending)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (h *participantHandler) PaymentConfirmation(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	username := fmt.Sprintf("%v", claims["username"])
	participant := role == string(entity.Participant)
	if participant {
		var req request.RequestTransactionID
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
		file, err := ctx.FormFile("receipt")
		if err != nil {
			errFormat := helper.ErrorFormatter(err)
			errMessage := helper.M{"error": errFormat}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusOK, response)
		}
		path := filepath.Base("/files/img/receipt/")
		filetype := strings.Split(file.Filename, ".")
		destFile := path + "trxID_" + strconv.Itoa(int(req.ID)) + "_receipt_" + username + filetype[1]
		if err := ctx.SaveUploadedFile(file, destFile); err != nil {
			errorFormatter := helper.ErrorFormatter(err)
			errorMessage := helper.M{"error": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Recipt has been sent.", nil)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}
