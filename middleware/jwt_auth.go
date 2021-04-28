package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(service services.JWTServices) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization neede.", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := service.ValidateToken(authHeader)
		if token.Valid {
			// claims := token.Claims.(jwt.MapClaims)
			// log.Println("Claim[user_id]: ", claims["user_id"])
			// log.Println("Claim[email]: ", claims["email"])
			// // log.Println("Claim[role]: ", claims["role"])
		} else {
			log.Println(err)
			response := helper.ResponseFormatter(http.StatusUnauthorized, "error", errors.New("token is not valid"), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
