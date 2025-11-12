package dto

import (
	"time"

	"github.com/google/uuid"
)

// UserResponse represents the structure of a user returned in API responses.
// @Description User response payload.
// @Name UserResponse
type UserResponse struct {
	ID             uuid.UUID `json:"id" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	FirstName      string    `json:"first_name" example:"John"`
	LastName       string    `json:"last_name" example:"Doe"`
	MiddleName     *string   `json:"middle_name,omitempty" example:"Michael"`
	Email          string    `json:"email" example:"user@mail.ru"`
	Phone          *string   `json:"phone,omitempty" example:"+77010000000"`
	KnowledgeLevel string    `json:"knowledge_level" example:"beginner"`
	Role           string    `json:"role" example:"tutor"`

	// Tutor-specific fields (optional)
	Rating            *float64 `json:"rating,omitempty" example:"4.8"`
	Portfolio         *string  `json:"portfolio,omitempty" example:"Experienced backend developer with 5+ years in Go."`
	TestimonialsCount *int     `json:"testimonials_count,omitempty" example:"12"`

	CreatedAt time.Time `json:"created_at" example:"2025-11-12T19:45:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-11-12T19:45:00Z"`
}
