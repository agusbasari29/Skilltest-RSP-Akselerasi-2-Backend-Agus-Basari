package request

import "github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"

type RequestTransaction struct {
	ParticipantId uint                 `json:"participant_id" binding:"required"`
	CreatorId     uint                 `json:"creator_id" binding:"required"`
	EventId       uint                 `json:"event_id" binding:"required"`
	Amount        float32              `json:"amount" binding:"required"`
	StatusPayment entity.StatusPayment `json:"status_payment" binding:"required"`
}
