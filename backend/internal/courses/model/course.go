package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
)

type CourseFormat string

const (
	CourseFormatGroup      CourseFormat = "group"
	CourseFormatIndividual CourseFormat = "individual"
	CourseFormatIntensive  CourseFormat = "intensive"
)

type Course struct {
	ID         uuid.UUID            `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string               `gorm:"type:varchar(200);not null"`
	Difficulty model.KnowledgeLevel `gorm:"type:varchar(20);not null"`
	Duration   string               `gorm:"type:varchar(100)"` // e.g. "12 weeks", "40h"
	Price      int64                `gorm:"not null"`
	Format     CourseFormat         `gorm:"type:varchar(20);not null"`
	AuthorID   uuid.UUID            `gorm:"type:uuid;not null;index"` // tutor user id
	Author     model.User           `gorm:"foreignKey:AuthorID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
