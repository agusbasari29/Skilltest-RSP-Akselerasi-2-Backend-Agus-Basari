package handler

import (
	"fmt"
	"net/http"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/response"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type creatorHandler struct {
	userServices services.UserServices
	jwtService   services.JWTServices
}

func NewCreatorHandler(userServices services.UserServices, jwtService services.JWTServices) *creatorHandler {
	return &creatorHandler{userServices, jwtService}
}

func (h *creatorHandler) CreateCreator(ctx *gin.Context) {
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
		var req request.RequestAuthRegister
		errReq := ctx.ShouldBind(&req)
		if errReq != nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		validationErr := validate.Struct(req)
		if validationErr != nil {
			errorFormatter := helper.ErrorFormatter(validationErr)
			errorMessage := helper.M{"error": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if h.userServices.UserIsExist(req.Username) {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User is alerady registered!", nil)
			ctx.AbortWithStatusJSON(http.StatusConflict, response)
			return
		}
		req.Role = entity.Creator
		newUser, err := h.userServices.CreateUser(req)
		if err != nil {
			errorFormatter := helper.ErrorFormatter(err)
			errorMessage := helper.M{"error": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		generatedToken := h.jwtService.GenerateToken(newUser)
		userData := response.ResponseUserFormatter(newUser)
		data := response.ResponseUserDataFormatter(userData, generatedToken)
		response := helper.ResponseFormatter(http.StatusOK, "success", "Creator sucessfully registered.", data)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *creatorHandler) UpdateCreator(ctx *gin.Context) {
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
		var req request.RequestUser
		errReq := ctx.ShouldBind(&req)
		if errReq != nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		validationErr := validate.Struct(req)
		if validationErr != nil {
			errorFormatter := helper.ErrorFormatter(validationErr)
			errorMessage := helper.M{"error": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if h.userServices.UserIsExist(req.Username) {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User is alerady registered!", nil)
			ctx.AbortWithStatusJSON(http.StatusConflict, response)
			return
		}
		req.Role = entity.Creator
		update, err := h.userServices.Update(req)
		if err != nil {
			errorFormatter := helper.ErrorFormatter(err)
			errorMessage := helper.M{"error": errorFormatter}
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Creator sucessfully updated.", update)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *creatorHandler) GetAllCreator(ctx *gin.Context) {
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
		var user entity.Users
		user.Role = entity.Creator
		result, err := h.userServices.GetAllCreator(user)
		if err != nil {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Failed to fetch all event.", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		var resultFormat []response.ResponseUser
		for _, v := range result {
			resultFormat = append(resultFormat, response.ResponseUserFormatter(v))
		}
		response := helper.ResponseFormatter(http.StatusOK, "success", "Successfully fetching data.", resultFormat)
		ctx.JSON(http.StatusOK, response)
	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *creatorHandler) DeleteCreator(ctx *gin.Context) {
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

	}
	response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User privilege...", nil)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}
