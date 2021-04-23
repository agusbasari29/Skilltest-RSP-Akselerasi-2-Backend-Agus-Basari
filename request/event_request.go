package request

import (
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
)

type RequestEvent struct {
	CreatorId         uint               `json:"creator_id" validate:"required"`
	TitleEvent        string             `json:"title_event" validate:"required"`
	LinkWebinar       string             `json:"link_webinar" validate:"required,url"`
	Description       string             `json:"description" validate:"required"`
	Banner            string             `json:"banner" validate:"required"`
	Price             float32            `json:"price" validate:"required,number"`
	Quantity          int                `json:"quantity" validate:"required,number"`
	Status            entity.EventStatus `json:"status" validate:"required"`
	EventStartDate    time.Time          `json:"event_start_date" validate:"required"`
	EventEndDate      time.Time          `json:"event_end_date" validate:"required"`
	CampaignStartDate time.Time          `json:"campaign_start_date" validate:"required"`
	CampaignEndDate   time.Time          `json:"campaign_end_date" validate:"required"`
}
