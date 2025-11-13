package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateScheduleRequest represents a schedule creation request.
// @Description Schedule creation request payload
// @Name CreateScheduleRequest
type CreateScheduleRequest struct {
	CourseID uuid.UUID  `json:"courseId" binding:"required" example:"f9c45610-e124-4d01-95cb-1c2a2c5c9999"`
	BranchID *uuid.UUID `json:"branchId,omitempty" example:"704e8400-e29b-41d4-a716-446655440000"`
	StartAt  time.Time  `json:"startAt" binding:"required" example:"2025-03-01T10:00:00Z"`
	EndAt    time.Time  `json:"endAt" binding:"required" example:"2025-04-01T14:00:00Z"`
	Capacity int        `json:"capacity" binding:"required,min=1" example:"25"`
}
