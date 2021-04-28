package repository

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"gorm.io/gorm"
)

type ReportRepository interface {
	DetailReportByEvent(eventId uint) ([]entity.Report, error)
	GetAllSummaryEvent() ([]entity.SummaryReport, error)
	GetAllSummaryEventByCreator(creatorId uint) ([]entity.SummaryReport, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *reportRepository {
	return &reportRepository{db}
}

func (r *reportRepository) DetailReportByEvent(eventId uint) ([]entity.Report, error) {
	var report []entity.Report
	err := r.db.Raw("SELECT tx.id, tx.event_id, e.title_event, e.description, u.fullname as creator, tx.amount as ticket_price, u1.id as participant_id, u1.fullname as name, u1.email 	FROM transactions tx JOIN events e ON tx.event_id = e.id JOIN users u ON e.creator_id = u.id JOIN users u1 ON tx.participant_id = u1.id WHERE tx.status_payment = 'passed' AND tx.event_id = ? 	GROUP BY tx.id, tx.event_id, e.title_event, e.description, creator, ticket_price, u1.id", eventId).Find(&report).Error
	if err != nil {
		return report, err
	}
	return report, nil
}

func (r *reportRepository) GetAllSummaryEvent() ([]entity.SummaryReport, error) {
	var summary []entity.SummaryReport
	err := r.db.Raw("SELECT e.title_event, e.description, tx.event_id, u.fullname as creator, sum(tx.amount) as total_amount, count(tx.event_id) as total_participant FROM transactions tx JOIN events e ON e.id = tx.event_id JOIN users u ON u.id = e.creator_id WHERE tx.status_payment = 'passed' GROUP BY e.id, tx.event_id, u.fullname ORDER BY total_amount DESC, total_participant DESC", nil).Find(&summary).Error
	if err != nil {
		return summary, err
	}
	return summary, nil
}

func (r *reportRepository) GetAllSummaryEventByCreator(creatorId uint) ([]entity.SummaryReport, error) {
	var summary []entity.SummaryReport
	err := r.db.Raw("SELECT e.title_event, e.description, tx.event_id, u.fullname as creator, sum(tx.amount) as total_amount, count(tx.event_id) as total_participant FROM transactions tx JOIN events e ON e.id = tx.event_id JOIN users u ON u.id = e.creator_id WHERE tx.status_payment = 'passed' AND tx.creator_id = ? GROUP BY e.id, tx.event_id, u.fullname ORDER BY total_amount DESC, total_participant DESC", creatorId).Find(&summary).Error
	if err != nil {
		return summary, err
	}
	return summary, nil
}
