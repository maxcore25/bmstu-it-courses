package dto

import (
	"github.com/google/uuid"
)

// UpdateCourseRequest represents the structure for updating a course.
type UpdateCourseRequest struct {
	Name       *string    `json:"name,omitempty" example:"Алгоритмы и структуры данных"`
	Difficulty *string    `json:"difficulty,omitempty" example:"beginner"`
	Duration   *string    `json:"duration,omitempty" example:"12 недель"`
	Price      *int64     `json:"price,omitempty" example:"20000"`
	Format     *string    `json:"format,omitempty" example:"group"`
	AuthorID   *uuid.UUID `json:"authorId,omitempty" example:"760e8400-e29b-41d4-a716-446655440000"`
}
