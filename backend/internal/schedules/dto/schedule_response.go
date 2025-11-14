package dto

import (
	"time"

	"github.com/google/uuid"
	branchDto "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/dto"
	courseDto "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/dto"
)

// ScheduleResponse represents the structure of a schedule returned in API responses.
// @Description Schedule response payload.
// @Name ScheduleResponse
type ScheduleResponse struct {
	ID        uuid.UUID  `json:"id" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	CourseID  uuid.UUID  `json:"courseId" example:"f9c45610-e124-4d01-95cb-1c2a2c5c9999"`
	BranchID  *uuid.UUID `json:"branchId,omitempty" example:"704e8400-e29b-41d4-a716-446655440000"`
	StartAt   time.Time  `json:"startAt" example:"2025-03-01T10:00:00Z"`
	EndAt     time.Time  `json:"endAt" example:"2025-04-01T14:00:00Z"`
	Capacity  int        `json:"capacity" example:"25"`
	Reserved  int        `json:"reserved" example:"12"`
	CreatedAt time.Time  `json:"createdAt" example:"2025-11-12T19:45:00Z"`
	UpdatedAt time.Time  `json:"updatedAt" example:"2025-11-13T19:45:00Z"`

	// Expandable fields
	Course *courseDto.CourseResponse `json:"course,omitempty"`
	Branch *branchDto.BranchResponse `json:"branch,omitempty"`
}
