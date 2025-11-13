package model

import (
	"time"

	"github.com/google/uuid"
)

type Branch struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Address   string    `gorm:"type:text;not null"`
	Rooms     int       `gorm:"not null"` // amount of rooms
	CreatedAt time.Time
	UpdatedAt time.Time
}
