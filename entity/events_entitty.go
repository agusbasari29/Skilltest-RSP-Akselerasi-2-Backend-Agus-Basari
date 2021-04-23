package entity

import (
	"time"

	"gorm.io/gorm"
)

type EventStatus string

const (
	Draft   EventStatus = "draft"
	Release EventStatus = "release"
)

type Event struct {
	gorm.Model
	ID                uint `gorm:"primariKey;autoIncrment"`
	CreatorId         int
	TitleEvent        string
	LinkWebinar       string
	Description       string
	Banner            string
	Price             float32
	Quantity          int
	Status            EventStatus
	EventStartDate    time.Time
	EventEndDate      time.Time
	CampaignStartDate time.Time
	CampaignEndDate   time.Time
}
