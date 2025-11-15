package dto

import (
	"github.com/google/uuid"
)

// CreateOrderRequest represents public input from API.
type CreateOrderInput struct {
	CourseID   uuid.UUID  `json:"courseId" example:"760e8400-e29b-41d4-a716-446655440002"`
	ScheduleID *uuid.UUID `json:"scheduleId,omitempty" example:"760e8400-e29b-41d4-a716-446655440002"`
	BranchID   *uuid.UUID `json:"branchId,omitempty" example:"760e8400-e29b-41d4-a716-446655440002"`
}

// CreateOrderRequest represents the structure for creating a new order.
type CreateOrderRequest struct {
	ClientID   uuid.UUID  `json:"clientId" example:"760e8400-e29b-41d4-a716-446655440002"`
	CourseID   uuid.UUID  `json:"courseId" example:"760e8400-e29b-41d4-a716-446655440001"`
	ScheduleID *uuid.UUID `json:"scheduleId,omitempty" example:"760e8400-e29b-41d4-a716-446655440003"`
	BranchID   *uuid.UUID `json:"branchId,omitempty" example:"760e8400-e29b-41d4-a716-446655440004"`
}
