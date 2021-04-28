package request

import (
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
)

type RequestEvent struct {
	CreatorId         int                `json:"creator_id"`
	TitleEvent        string             `json:"title_event"`
	LinkWebinar       string             `json:"link_webinar"`
	Description       string             `json:"description"`
	Banner            string             `json:"banner"`
	Price             float32            `json:"price"`
	Quantity          int                `json:"quantity"`
	Status            entity.EventStatus `json:"status"`
	EventStartDate    time.Time          `json:"event_start_date"`
	EventEndDate      time.Time          `json:"event_end_date"`
	CampaignStartDate time.Time          `json:"campaign_start_date"`
	CampaignEndDate   time.Time          `json:"campaign_end_date"`
}

type RequestEventByID struct {
	ID uint `json:"id" vaildate:"required"`
}
