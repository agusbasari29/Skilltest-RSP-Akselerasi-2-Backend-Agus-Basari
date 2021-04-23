package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/response"
	"github.com/dgrijalva/jwt-go"
)

type JWTServices interface {
	GenerateToken(user entity.Users) response.ResponseCredential
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secret string
	issuer string
}

type jwtCustomClaim struct {
	UserID uint            `json:"user_id"`
	Role   entity.UserRole `json:"role"`
	Email  string          `json:"email"`
	jwt.StandardClaims
}

func NewJWTService() JWTServices {
	return &jwtService{
		issuer: "xjx",
		secret: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET KEY")
	if secretKey == "" {
		secretKey = "jabrix"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(user entity.Users) response.ResponseCredential {
	claims := &jwtCustomClaim{}
	claims.UserID = user.ID
	claims.Role = user.Role
	claims.Email = user.Email
	claims.ExpiresAt = time.Now().AddDate(1, 0, 0).Unix()
	claims.Issuer = j.issuer
	claims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secret))
	if err != nil {
		panic(err)
	}
	credential := response.ResponseCredential{}
	credential.UserID = claims.UserID
	credential.Token = t
	credential.Issuer = claims.Issuer
	credential.IssuedAt = claims.IssuedAt
	credential.ExpiresAt = claims.ExpiresAt
	credential.Role = claims.Role

	return credential
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	jwtString := strings.Split(token, "Bearer ")[1]
	return jwt.Parse(jwtString, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}
