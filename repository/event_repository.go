package repository

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	InsertEvent(event entity.Event) (entity.Event, error)
	GetAllEvent() []entity.Event
	UpdateEvent(event entity.Event) (entity.Event, error)
	DeleteEvent(event entity.Event) error
	GetEventByStatus(even entity.Event) []entity.Event
	GetEventByID(event entity.Event) (*entity.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *eventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) InsertEvent(event entity.Event) (entity.Event, error) {
	err := r.db.Raw("INSERT INTO event (creator_id, title_event, link_webinar,description, banner, price, quantity, status, event_start_date, event_end_date, campaign_start_date, campaign_end_date) VALUES (@CreatorId, @TitleEvent, @LinkWebinar, @Description, @Banner, @Price, @Quantity, @Status, @EventStartDate, @EventEndDate, @CampaignStartDate, @CampaignEndDate)", event).Save(&event).Error
	if err != nil {
		return event, err
	}
	return event, nil
}

func (r *eventRepository) GetAllEvent() []entity.Event {
	var event []entity.Event
	result := r.db.Raw("SELECT * FROM event").Find(&event)
	if result != nil {
		return event
	}
	return nil
}

func (r *eventRepository) UpdateEvent(event entity.Event) (entity.Event, error) {
	err := r.db.Raw("UPDATE event SET creator_id = @CreatorId, title_event = @TitleEvent, link_webinar = @LinkWebinar, description = @Description, banner = @Banner, price = @Price, quantity = @Quantity, status = @Status, event_start_date = @EventStartDate, event_end_date = @EventEndDate, campaign_start_date = @CampaignStartDate, campaign_end_date = @CampaignEndDate WHERE id = @ID", event).Save(&event).Error
	if err != nil {
		return event, err
	}
	return event, nil
}

func (r *eventRepository) DeleteEvent(event entity.Event) error {
	err := r.db.Raw("UPDATE event SET deleted_at = &DeletedAt WHERE id = @ID", event).Save(&event).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *eventRepository) GetEventByStatus(event entity.Event) []entity.Event {
	var events []entity.Event
	err := r.db.Raw("SELECT * FROM event WHERE status = @Status", event).Find(&events).Error
	if err != nil {
		return nil
	}
	return events
}

func (r *eventRepository) GetEventByID(event entity.Event) (*entity.Event, error) {
	err := r.db.Raw("SELECT * FROM event WHERE id = @ID").Take(&event).Error
	if err != nil {
		return &event, err
	}
	return &event, nil
}
