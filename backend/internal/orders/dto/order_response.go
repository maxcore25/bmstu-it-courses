package dto

import (
	"github.com/google/uuid"
	userDto "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/dto"
	branchDto "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/dto"
	courseDto "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/dto"
)

// OrderResponse represents the structure for responding with order data.
type OrderResponse struct {
	ID         uuid.UUID  `json:"id" example:"760e8400-e29b-41d4-a716-446655440000"`
	ClientID   uuid.UUID  `json:"clientId" example:"760e8400-e29b-41d4-a716-446655440002"`
	CourseID   uuid.UUID  `json:"courseId" example:"760e8400-e29b-41d4-a716-446655440001"`
	ScheduleID *uuid.UUID `json:"scheduleId,omitempty" example:"760e8400-e29b-41d4-a716-446655440003"`
	BranchID   *uuid.UUID `json:"branchId,omitempty" example:"760e8400-e29b-41d4-a716-446655440004"`
	Price      int64      `json:"price" example:"99900"`
	CreatedAt  string     `json:"createdAt" example:"2024-04-10T15:04:05Z"`
	UpdatedAt  string     `json:"updatedAt" example:"2024-04-10T15:04:05Z"`

	// Expandable fields
	Client *userDto.UserResponse     `json:"client,omitempty"`
	Course *courseDto.CourseResponse `json:"course,omitempty"`
	Branch *branchDto.BranchResponse `json:"branch,omitempty"`
}
