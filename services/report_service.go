package services

import (
	"log"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/response"
)

type ReportServices interface {
	GetReportByEvent(eventId uint) response.ResponseReport
	GetAllSummaryEvent() []entity.SummaryReport
	GetAllSummaryEventByCreator(creatorId uint) []entity.SummaryReport
}

type reportServices struct {
	reportRepository repository.ReportRepository
}

func NewReportService(reportRepository repository.ReportRepository) *reportServices {
	return &reportServices{reportRepository}
}

func (s *reportServices) GetReportByEvent(eventId uint) response.ResponseReport {
	var resp response.ResponseReport
	var participant []response.EventParticipant
	report, err := s.reportRepository.DetailReportByEvent(eventId)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, v := range report {
		var pars response.EventParticipant
		pars.ParticipantId = v.ParticipantId
		pars.Name = v.Name
		pars.Email = v.Email
		participant = append(participant, pars)
	}
	resp.ID = report[0].ID
	resp.TitleEvent = report[0].TitleEvent
	resp.Description = report[0].Description
	resp.Creator = report[0].Creator
	resp.TicketPrice = report[0].TicketPrice
	resp.TotalParticipant = len(report)
	resp.TotalAmount = float32(len(report)) * report[0].TicketPrice
	resp.Participant = participant
	return resp
}

func (s *reportServices) GetAllSummaryEvent() []entity.SummaryReport {
	summary, err := s.reportRepository.GetAllSummaryEvent()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return summary
}

func (s *reportServices) GetAllSummaryEventByCreator(creatorId uint) []entity.SummaryReport {
	summary, err := s.reportRepository.GetAllSummaryEventByCreator(creatorId)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return summary
}
