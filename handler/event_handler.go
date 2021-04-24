package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/response"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	eventServices       services.EventServices
	transactionServices services.TransactionServices
	jwtService          services.JWTServices
}

func NewEventHandler(eventServices services.EventServices, jwtService services.JWTServices, transactionServices services.TransactionServices) *eventHandler {
	return &eventHandler{eventServices, transactionServices, jwtService}
}

func (h *eventHandler) CreateEvent(ctx *gin.Context) {
	var req request.RequestEvent
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	admin := role == string(entity.Admin)
	if admin || role == string(entity.Creator) {
		err = ctx.ShouldBind(&req)
		if err != nil {
			errorMessage := helper.M{"error": err.Error()}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		err = validate.Struct(req)
		if err != nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid input type", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		id, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
		if err != nil {
			log.Fatalf(err.Error())
		}
		req.CreatorId = id
		newEvent, err := h.eventServices.CreateEvent(req)
		if err != nil {
			errorFormatter := helper.ErrorFormatter(err)
			errorMessage := helper.M{"error": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		eventFormat := response.ResponseEventFormatter(newEvent)
		response := helper.ResponseFormatter(http.StatusOK, "success", "New event successfully created.", eventFormat)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *eventHandler) UpdateEvent(ctx *gin.Context) {
	var req request.RequestEvent
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	admin := role == string(entity.Admin)
	if admin || role == string(entity.Creator) {
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
		updateEvent, err := h.eventServices.UpdateEvent(req)
		if err != nil {
			errorFormatter := helper.ErrorFormatter(err)
			errorMessage := helper.M{"error": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		eventFormat := response.ResponseEventFormatter(updateEvent)
		response := helper.ResponseFormatter(http.StatusOK, "success", "Event successfully updated.", eventFormat)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *eventHandler) GetAllEvent(ctx *gin.Context) {
	var req request.RequestEvent
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	admin := role == string(entity.Admin)
	if admin || role == string(entity.Creator) {
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
		events := h.eventServices.GetAllEvent()
		if events == nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to fetch all event.", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		var eventFormat []response.ResponseEvent
		for _, v := range events {
			eventFormat = append(eventFormat, response.ResponseEventFormatter(v))
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully fetching data.", eventFormat)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *eventHandler) DeletedEvent(ctx *gin.Context) {
	var req request.RequestEventByID
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
		delete := h.eventServices.DeleteEvent(req)
		if !delete {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to delete event.", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Event successfully deleted.", nil)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *eventHandler) GetEventByReleaseStatus(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	if role == string(entity.Participant) {
		release := h.eventServices.GetReleaseEvent()
		if release == nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to fetch all release event data.", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		var eventFormat []response.ResponseEvent
		for _, v := range release {
			eventFormat = append(eventFormat, response.ResponseEventFormatter(v))
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully fetching all release event data.", eventFormat)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *eventHandler) MakeEventPurchase(ctx *gin.Context) {
	var req request.RequestEventByID
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		panic(err.Error())
	}
	role := fmt.Sprintf("%v", claims["role"])
	if role == string(entity.Participant) {
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
		event, err := h.eventServices.GetEventByID(req)
		if err != nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to retrieve data event", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		var reqTrx request.RequestTransaction
		reqTrx.EventId = event.ID
		totalTrx, err := h.transactionServices.GetTransactionByEventID(reqTrx)
		if err != nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to get total per event transaction", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		if len(totalTrx) > event.Quantity {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Event sold out!", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		var trx entity.Transaction
		trx.EventId = int(event.ID)
		trx.ParticipantId = id
		trx.StatusPayment = entity.Passed
		oneBuyOne := h.transactionServices.GetTransactionByEventAndParticipantAndStatusPayment(trx)
		if oneBuyOne {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Transacton is exist!", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		trx.Amount = event.Price
		var newTrx request.RequestTransaction
		createTrx, err := h.transactionServices.CreateTransaction(newTrx)
		if err != nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to create a new transaction.", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		formatter := response.ResponseTransactionFormatter(createTrx)
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully create new transaction.", formatter)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}
