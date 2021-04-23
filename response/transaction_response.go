package response

import (
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"gorm.io/gorm"
)

type ResponseTransaction struct {
	ID            uint                 `json:"id"`
	ParticipantId int                  `json:"participant_id"`
	CreatorId     int                  `json:"creator_id"`
	EventId       int                  `json:"event_id"`
	Amount        float32              `json:"amount"`
	StatusPayment entity.StatusPayment `json:"status_payment"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	DeletedAt     gorm.DeletedAt       `json:"deleted_at"`
}

func ResponseTransactionFormatter(trans entity.Transaction) ResponseTransaction {
	formatter := ResponseTransaction{}
	formatter.ID = trans.ID
	formatter.ParticipantId = trans.ParticipantId
	formatter.CreatorId = trans.CreatorId
	formatter.EventId = trans.EventId
	formatter.Amount = trans.Amount
	formatter.StatusPayment = trans.StatusPayment
	formatter.CreatedAt = trans.CreatedAt
	formatter.UpdatedAt = trans.UpdatedAt
	formatter.DeletedAt = trans.DeletedAt

	return formatter
}
