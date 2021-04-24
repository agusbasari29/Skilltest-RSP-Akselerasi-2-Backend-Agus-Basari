package services

import (
	"log"
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/mashingan/smapping"
)

type EventServices interface {
	CreateEvent(req request.RequestEvent) (entity.Event, error)
	UpdateEvent(req request.RequestEvent) (entity.Event, error)
	GetAllEvent() []entity.Event
	DeleteEvent(req request.RequestEventByID) bool
	GetReleaseEvent() []entity.Event
	GetEventByID(req request.RequestEventByID) (*entity.Event, error)
}

type eventServices struct {
	eventRepository repository.EventRepository
}

func NewEventServices(eventRepository repository.EventRepository) *eventServices {
	return &eventServices{eventRepository}
}

func (s *eventServices) CreateEvent(req request.RequestEvent) (entity.Event, error) {
	var event entity.Event
	err := smapping.FillStruct(&event, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	event.CreatedAt = time.Now()
	createdEvent, err := s.eventRepository.InsertEvent(event)
	if err != nil {
		return event, err
	}
	return createdEvent, nil
}

func (s *eventServices) UpdateEvent(req request.RequestEvent) (entity.Event, error) {
	var event entity.Event
	err := smapping.FillStruct(&event, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	event.UpdatedAt = time.Now()
	updatedEvent, err := s.eventRepository.UpdateEvent(event)
	if err != nil {
		return event, err
	}
	return updatedEvent, nil
}

func (s *eventServices) GetAllEvent() []entity.Event {
	result := s.eventRepository.GetAllEvent()
	if result != nil {
		return result
	}
	return nil
}

func (s *eventServices) DeleteEvent(req request.RequestEventByID) bool {
	var event entity.Event
	event.ID = req.ID
	err := s.eventRepository.DeleteEvent(event)
	return err == nil
}

func (s *eventServices) GetReleaseEvent() []entity.Event {
	var event entity.Event
	event.Status = entity.Release
	result := s.eventRepository.GetEventByStatus(event)
	if result == nil {
		return nil
	}
	return result
}

func (s *eventServices) GetEventByID(req request.RequestEventByID) (*entity.Event, error) {
	var event entity.Event
	event.ID = req.ID
	result, err := s.eventRepository.GetEventByID(event)
	if err != nil {
		return &event, err
	}
	return result, nil
}
