package handler

import (
	"net/http"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/response"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	userServices services.UserServices
	jwtService   services.JWTServices
}

func NewAuthHandler(userServices services.UserServices, jwtService services.JWTServices) *authHandler {
	return &authHandler{userServices, jwtService}
}

func (h *authHandler) Register(c *gin.Context) {
	var req request.RequestAuthRegister
	errReq := c.ShouldBind(&req)
	if errReq != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	validationErr := validate.Struct(req)
	if validationErr != nil {
		errorFormatter := helper.ErrorFormatter(validationErr)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if h.userServices.UserIsExist(req.Username) {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User is alerady registered!", nil)
		c.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}
	newUser, err := h.userServices.CreateUser(req)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	generatedToken := h.jwtService.GenerateToken(newUser)
	userData := response.ResponseUserFormatter(newUser)
	data := response.ResponseUserDataFormatter(userData, generatedToken)
	response := helper.ResponseFormatter(http.StatusOK, "success", "User sucessfully registered.", data)
	c.JSON(http.StatusOK, response)

}

func (h *authHandler) Login(c *gin.Context) {
	var req request.RequestAuthLogin

	err := c.ShouldBind(&req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	validationErr := validate.Struct(req)
	if validationErr != nil {
		errorFormatter := helper.ErrorFormatter(validationErr)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	credential := h.userServices.VerifyCredential(req.Username, req.Password)
	if v, ok := credential.(entity.Users); ok {
		generatedToken := h.jwtService.GenerateToken(v)
		userData := response.ResponseUserFormatter(v)
		data := response.ResponseUserDataFormatter(userData, generatedToken)
		response := helper.ResponseFormatter(http.StatusOK, "success", "User sucessfully registered.", data)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusUnauthorized, "error", "Cannot log in!", nil)
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (h *authHandler) ForgetPassword(c *gin.Context) {
	var req request.RequestAuthForgetPassword
	err := c.ShouldBind(&req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	validationErr := validate.Struct(req)
	if validationErr != nil {
		errorFormatter := helper.ErrorFormatter(validationErr)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	EmailIsExist := h.userServices.EmailIsExist(req)
	if !EmailIsExist {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "User with this email is not registered.", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

}
