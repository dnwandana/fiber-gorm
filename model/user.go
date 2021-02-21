package model

import (
	"time"
)

// User struct representation of database model
type User struct {
	ID        uint      `gorm:"primarykey"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
