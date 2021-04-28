package entity

type Report struct {
	ID            uint    `json:"id"`
	TitleEvent    string  `json:"title_event"`
	Description   string  `json:"description"`
	Creator       string  `json:"creator"`
	TicketPrice   float32 `json:"ticket_price"`
	ParticipantId uint    `json:"participant_id"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
}

type SummaryReport struct {
	EventId          uint    `json:"event_id"`
	TitleEvent       string  `json:"title_event"`
	Description      string  `json:"description"`
	Creator          string  `json:"creator"`
	TotalAmount      float32 `json:"total_amount"`
	TotalParticipant int     `json:"total_participant"`
}
