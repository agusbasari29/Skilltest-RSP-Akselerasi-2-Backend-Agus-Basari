package repository

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"gorm.io/gorm"
)

type ReportRepository interface {
}

type reportRepository struct {
	db *gorm.DB
}

type ReportByEvent struct {
	TitleEvent       string `json:"title_event"`
	Description      string `json:"description"`
	TotalParticipant int    `json:"total_participant"`
}

func NewReportRepository(db *gorm.DB) *reportRepository {
	return &reportRepository{db}
}

func (r *reportRepository) DetailReportByEvent(event entity.Event) ([]ReportByEvent, error) {
	var report []ReportByEvent
	err := r.db.Raw("SELECT e.title_event e.decsription  t.amount sum(t.amount) as total_amount u.fullname FROM events e INNER JOIN  transactions t on e.id = t.event_id ", event).Find(&report).Error
	if err != nil {
		return report, err
	}
	return report, nil
}
