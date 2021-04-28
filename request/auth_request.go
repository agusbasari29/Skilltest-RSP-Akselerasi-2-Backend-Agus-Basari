package request

import "github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"

type RequestAuthRegister struct {
	Username string          `json:"username" validate:"required,alphanum"`
	Fullname string          `json:"fullname" validate:"required"`
	Email    string          `json:"email" validate:"required,email"`
	Password string          `json:"password" validate:"required"`
	Role     entity.UserRole `json:"role"`
}

type RequestAuthLogin struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required"`
}

type RequestAuthUpdate struct {
	ID       uint            `json:"id"`
	Username string          `json:"username"`
	Fullname string          `json:"fullname"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Role     entity.UserRole `json:"role"`
}


type RequestAuthForgetPassword struct {
	Email string `json:"email" validate:"required,email"`
}
