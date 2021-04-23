package request

import "github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"

type RequestUser struct {
	ID       uint            `json:"id"`
	Username string          `json:"username" validate:"required,alphanum"`
	Fullname string          `json:"fullname" validate:"required"`
	Email    string          `json:"email" validate:"required,email"`
	Password string          `json:"password" validate:"required"`
	Role     entity.UserRole `json:"role"`
}

type RequestUserProfile struct {
	ID uint `json:"id"`
}
