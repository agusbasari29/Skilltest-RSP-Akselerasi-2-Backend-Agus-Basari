package response

import (
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"gorm.io/gorm"
)

type ResponseUserData struct {
	User       interface{} `json:"user_data"`
	Credential interface{} `json:"credential"`
}

type ResponseUser struct {
	ID             uint            `json:"id"`
	Username       string          `json:"username"`
	Fullname       string          `json:"fullname"`
	Email          string          `json:"email"`
	Role           entity.UserRole `json:"role"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      gorm.DeletedAt  `json:"deleted_at"`
	UserCredential interface{}     `json:"user_credential"`
}

func ResponseUserFormatter(user entity.Users) ResponseUser {

	formatter := ResponseUser{}
	formatter.ID = user.ID
	formatter.Username = user.Username
	formatter.Fullname = user.Fullname
	formatter.Email = user.Email
	formatter.Role = user.Role
	formatter.CreatedAt = user.CreatedAt
	formatter.UpdatedAt = user.UpdatedAt
	formatter.DeletedAt = user.DeletedAt

	return formatter
}

func ResponseUserDataFormatter(user interface{}, credential interface{}) ResponseUserData {
	userData := ResponseUserData{
		User:       user,
		Credential: credential,
	}
	return userData
}
