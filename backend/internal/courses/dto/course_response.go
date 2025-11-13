package dto

import (
	"time"

	"github.com/google/uuid"
)

// CourseResponse represents the response structure for a Course.
type CourseResponse struct {
	ID         uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name       string    `json:"name" example:"Алгоритмы и структуры данных"`
	Difficulty string    `json:"difficulty" example:"beginner"`
	Duration   string    `json:"duration" example:"12 weeks"`
	Price      int64     `json:"price" example:"20000"`
	Format     string    `json:"format" example:"group"`
	AuthorID   uuid.UUID `json:"authorId" example:"760e8400-e29b-41d4-a716-446655440000"`
	CreatedAt  time.Time `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	UpdatedAt  time.Time `json:"updatedAt" example:"2024-01-05T00:00:00Z"`
}
