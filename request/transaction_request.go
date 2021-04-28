package request

import "github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"

type RequestTransaction struct {
	ParticipantId uint                 `json:"participant_id" validate:"required"`
	CreatorId     uint                 `json:"creator_id" validate:"required"`
	EventId       uint                 `json:"event_id" validate:"required"`
	Amount        float32              `json:"amount" validate:"required"`
	StatusPayment entity.StatusPayment `json:"status_payment"`
}

type RequestAllParticipantTransaction struct {
	ID            uint                 `json:"id" validate:"required"`
	StatusPayment entity.StatusPayment `json:"status_payment" validate:"required"`
}

type RequestParticipantTransaction struct {
	ID            uint                 `json:"id" validate:"required"`
	ParticipantId uint                 `json:"participant_id" validate:"required"`
	StatusPayment entity.StatusPayment `json:"status_payment" validate:"required"`
}

type RequestTransactionUpdate struct {
	ID            uint    `json:"id" validate:"required"`
	ParticipantId uint    `json:"participant_id" validate:"required"`
	CreatorId     uint    `json:"creator_id" validate:"required"`
	EventId       uint    `json:"event_id" validate:"required"`
	Amount        float32 `json:"amount" validate:"required"`
	Receipt       string
	StatusPayment entity.StatusPayment `json:"status_payment"`
}

type RequestTransactionID struct {
	ID uint `json:"id" validate:"required"`
}
