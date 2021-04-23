package entity

import "gorm.io/gorm"

type UserRole string

const (
	Admin       UserRole = "admin"
	Creator     UserRole = "creator"
	Participant UserRole = "participant"
)

type Users struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Fullname string
	Email    string
	Password string
	Role     UserRole
}
