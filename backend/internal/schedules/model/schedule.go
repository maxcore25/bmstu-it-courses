package model

import (
	"time"

	"github.com/google/uuid"
	branchModel "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/model"
	courseModel "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"
)

type Schedule struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CourseID  uuid.UUID `gorm:"type:uuid;not null;index"`
	Course    courseModel.Course
	BranchID  *uuid.UUID `gorm:"type:uuid;index"` // nil for online
	Branch    *branchModel.Branch
	StartAt   time.Time `gorm:"not null"`
	EndAt     time.Time `gorm:"not null"`
	Capacity  int       `gorm:"not null"`           // total seats for this schedule
	Reserved  int       `gorm:"not null;default:0"` // how many seats reserved
	CreatedAt time.Time
	UpdatedAt time.Time
}
