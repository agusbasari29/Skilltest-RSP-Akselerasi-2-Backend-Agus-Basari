package response

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
)

type ResponseCredential struct {
	Token     string          `json:"token"`
	UserID    uint            `json:"user_id"`
	Role      entity.UserRole `json:"role"`
	Issuer    string          `json:"issuer"`
	IssuedAt  int64           `json:"issued_at"`
	ExpiresAt int64           `json:"expired_at"`
}
