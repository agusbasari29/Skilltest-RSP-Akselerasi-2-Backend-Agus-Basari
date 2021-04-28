package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type reportHandler struct {
	reportServices services.ReportServices
	jwtServices    services.JWTServices
}

func NewReportHandler(reportServices services.ReportServices, jwtServices services.JWTServices) *reportHandler {
	return &reportHandler{reportServices, jwtServices}
}

func (h *reportHandler) DetailReportByEvent(ctx *gin.Context) {
	var req request.RequestEventByID
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtServices.ValidateToken(authHeader)
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
			log.Fatalf("%+v", err)
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
		resp := h.reportServices.GetReportByEvent(req.ID)
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully fetching report data", resp)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (h *reportHandler) AllSummaryReport(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtServices.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])
	admin := role == string(entity.Admin)
	if admin {
		resp := h.reportServices.GetAllSummaryEvent()
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully fetching report data", resp)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}

func (h *reportHandler) AllSummaryReportByCreator(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := h.jwtServices.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	role := fmt.Sprintf("%v", claims["role"])
	creator := role == string(entity.Creator)
	if creator {
		resp := h.reportServices.GetAllSummaryEventByCreator(uint(id))
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully fetching report data", resp)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
}
