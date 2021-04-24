package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type userHandler struct {
	userService services.UserServices
	jwtService  services.JWTServices
}

func NewUserHandler(userService services.UserServices, jwtService services.JWTServices) *userHandler {
	return &userHandler{userService, jwtService}
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var req request.RequestUser
	err := c.ShouldBind(req)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "data_type", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	valErr := validate.Struct(req)
	if valErr != nil {
		errorFormatter := helper.ErrorFormatter(valErr)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "validation", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseFormatter(http.StatusOK, "OK", "OK", nil)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Profile(c *gin.Context) {
	var req request.RequestUserProfile
	authHeader := c.GetHeader("Authorization")
	token, err := h.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	req.ID = uint(id)
	user, err := h.userService.Profile(req)
	if err != nil {
		panic(err.Error())
	}

	response := helper.ResponseFormatter(http.StatusOK, "OK", "OK", user)
	c.JSON(http.StatusOK, response)
}
