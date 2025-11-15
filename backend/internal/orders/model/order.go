package model

import (
	"time"

	"github.com/google/uuid"
	authModel "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	branchModel "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/model"
	courseModel "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"
	scheduleModel "github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/model"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ClientID   uuid.UUID `gorm:"type:uuid;not null;index"` // user id of client
	Client     authModel.User
	CourseID   uuid.UUID `gorm:"type:uuid;not null;index"`
	Course     courseModel.Course
	ScheduleID *uuid.UUID `gorm:"type:uuid;index"` // optional, if ordering for a specific schedule
	Schedule   *scheduleModel.Schedule
	BranchID   *uuid.UUID `gorm:"type:uuid;index"` // optional duplicate for convenience
	Branch     *branchModel.Branch
	Price      int64 `gorm:"not null"` // snapshot of price at time of order
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
