package response

import (
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"gorm.io/gorm"
)

type ResponseEvent struct {
	ID                uint               `json:"id"`
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
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	DeletedAt         gorm.DeletedAt     `json:"deleted_at"`
}

func ResponseEventFormatter(event entity.Event) ResponseEvent {
	formatter := ResponseEvent{}
	formatter.ID = event.ID
	formatter.CreatorId = event.CreatorId
	formatter.TitleEvent = event.TitleEvent
	formatter.LinkWebinar = event.LinkWebinar
	formatter.Description = event.Description
	formatter.Banner = event.Banner
	formatter.Price = event.Price
	formatter.Quantity = event.Quantity
	formatter.Status = event.Status
	formatter.EventStartDate = event.EventStartDate
	formatter.EventEndDate = event.EventEndDate
	formatter.CampaignStartDate = event.CampaignStartDate
	formatter.CampaignEndDate = event.CampaignEndDate
	formatter.CreatedAt = event.CreatedAt
	formatter.UpdatedAt = event.UpdatedAt
	formatter.DeletedAt = event.DeletedAt
	return formatter
}
