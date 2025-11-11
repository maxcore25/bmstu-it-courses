package model

import (
	"time"

	"github.com/google/uuid"
)

type KnowledgeLevel string

const (
	LevelBeginner     KnowledgeLevel = "beginner"
	LevelIntermediate KnowledgeLevel = "intermediate"
	LevelAdvanced     KnowledgeLevel = "advanced"
)

type User struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName      string         `gorm:"type:varchar(50);not null"`
	LastName       string         `gorm:"type:varchar(50);not null"`
	MiddleName     *string        `gorm:"type:varchar(50)"`
	Email          string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password       string         `gorm:"type:varchar(255);not null"` // TODO: hash before save
	Phone          *string        `gorm:"type:varchar(20)"`
	KnowledgeLevel KnowledgeLevel `gorm:"type:varchar(20);not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
