package response

type ResponseReport struct {
	ID               uint    `json:"id"`
	TitleEvent       string  `json:"title_event"`
	Description      string  `json:"description"`
	Creator          string  `json:"creator"`
	TicketPrice      float32 `json:"ticket_price"`
	TotalParticipant int     `json:"total_participant"`
	TotalAmount      float32 `json:"total_amount"`
	Participant      []EventParticipant
}

type EventParticipant struct {
	ParticipantId uint   `json:"participant_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
}
