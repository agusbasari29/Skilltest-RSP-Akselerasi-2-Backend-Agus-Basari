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
	DeleteEvent(eventId uint) bool
	GetReleaseEvent() []entity.Event
	GetEventByID(eventId uint) (*entity.Event, error)
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

func (s *eventServices) DeleteEvent(eventId uint) bool {
	err := s.eventRepository.DeleteEvent(eventId)
	return err == nil
}

func (s *eventServices) GetReleaseEvent() []entity.Event {
	status := "release"
	result := s.eventRepository.GetEventByStatus(status)
	if result == nil {
		return nil
	}
	return result
}

func (s *eventServices) GetEventByID(eventId uint) (*entity.Event, error) {
	result, err := s.eventRepository.GetEventByID(eventId)
	if err != nil {
		return result, err
	}
	return result, nil
}
