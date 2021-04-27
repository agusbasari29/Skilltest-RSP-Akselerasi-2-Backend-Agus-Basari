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
	err := r.db.Raw("SELECT e.id, e.title_event, e.description, ctr.fullname as creator, tx.amount, sum(tx.amount) as total_amount, u.fullname as participant FROM events as e LEFT JOIN users as ctr ON e.creator_id = ctr.id INNER JOIN transactions as tx ON e.id = tx.event_id INNER JOIN users as u ON tx.participant_id = u.id WHERE e.id = @ID GROUP BY e.id, tx.amount, e.title_event, ctr.id, participant, e.description ORDER BY e.id ASC", event).Find(&report).Error
	if err != nil {
		return report, err
	}
	return report, nil
}
