package handler

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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
	jwtService          services.JWTServices
	transactionServices services.TransactionServices
}

func NewEventHandler(eventServices services.EventServices, jwtService services.JWTServices, transactionServices services.TransactionServices) *eventHandler {
	return &eventHandler{eventServices, jwtService, transactionServices}
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
	creator := role == string(entity.Creator)
	if admin || creator {
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
		file, err := ctx.FormFile("upload")
		if err != nil {
			errFormat := helper.ErrorFormatter(err)
			errMessage := helper.M{"error": errFormat}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusOK, response)
		}
		path := filepath.Base("/files/img/banner/")
		filetype := strings.Split(file.Filename, ".")
		destFile := path + req.TitleEvent + filetype[1]
		if err := ctx.SaveUploadedFile(file, destFile); err != nil {
			errorFormatter := helper.ErrorFormatter(err)
			errorMessage := helper.M{"error": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		req.Banner = destFile
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
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
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
	creator := role == string(entity.Creator)
	if admin || creator {
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
		file, err := ctx.FormFile("upload")
		if err != nil {
			errFormat := helper.ErrorFormatter(err)
			errMessage := helper.M{"error": errFormat}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusOK, response)
		}
		path := filepath.Base("/files/img/banner/")
		filetype := strings.Split(file.Filename, ".")
		destFile := path + req.TitleEvent + filetype[1]
		if err := ctx.SaveUploadedFile(file, destFile); err != nil {
			errorFormatter := helper.ErrorFormatter(err)
			errorMessage := helper.M{"error": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		req.Banner = destFile
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
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
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
	creator := role == string(entity.Creator)
	if admin || creator {
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
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (h *eventHandler) DeleteEvent(ctx *gin.Context) {
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
		eventId, _ := strconv.Atoi(ctx.Param("id"))
		delete := h.eventServices.DeleteEvent(uint(eventId))
		if !delete {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to delete event.", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Event successfully deleted.", nil)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
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
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (h *eventHandler) GetEventDetail(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	if role == string(entity.Participant) {
		eventId, _ := strconv.Atoi(ctx.Param("id"))
		// e := entity.{ID: req.ID}
		// id := strconv.Itoa(int(e.ID))
		var event *entity.Event
		if event == nil {
			event, err := h.eventServices.GetEventByID(uint(eventId))
			if err != nil {
				response := helper.ResponseFormatter(http.StatusBadRequest, "error", "No event found.", nil)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
				return
			}
			response := helper.ResponseFormatter(http.StatusOK, "success", "Success retreiving event detail", event)
			ctx.JSON(http.StatusOK, response)
		} else {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to retreiving event detail.", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (h *eventHandler) MakeEventPurchase(ctx *gin.Context) {
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
		eventId, _ := strconv.Atoi(ctx.Param("id"))
		event, err := h.eventServices.GetEventByID(uint(eventId))
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
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Event ticket is sold out!", nil)
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
		//TODO notif
		formatter := response.ResponseTransactionFormatter(createTrx)
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully create new transaction.", formatter)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}
