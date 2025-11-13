package dto

import (
	"github.com/google/uuid"
)

// CreateCourseRequest represents the structure for creating a new course.
type CreateCourseRequest struct {
	Name       string    `json:"name" example:"Algorithms and Data Structures"`
	Difficulty string    `json:"difficulty" example:"beginner"`
	Duration   string    `json:"duration" example:"12 weeks"`
	Price      int64     `json:"price" example:"20000"`
	Format     string    `json:"format" example:"group"`
	AuthorID   uuid.UUID `json:"authorId" example:"760e8400-e29b-41d4-a716-446655440000"`
}
