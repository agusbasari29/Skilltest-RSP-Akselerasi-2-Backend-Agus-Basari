package entity

import "gorm.io/gorm"

type StatusPayment string

const (
	Passed StatusPayment = "passed"
	Failed StatusPayment = "failed"
)

type Transaction struct {
	gorm.Model
	ID            uint `gorm:"primaryKey;autoIncrement"`
	ParticipantId int
	Participant   Users `gorm:"foreignKey:ParticipantId"`
	CreatorId     int
	EventId       int
	Amount        float32
	StatusPayment StatusPayment
}
