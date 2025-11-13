package mapper

import (
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"
)

// NewCourseResponse creates a CourseResponse from model.Course.
func NewCourseResponse(c *model.Course) *dto.CourseResponse {
	return &dto.CourseResponse{
		ID:         c.ID,
		Name:       c.Name,
		Difficulty: string(c.Difficulty),
		Duration:   c.Duration,
		Price:      c.Price,
		Format:     string(c.Format),
		AuthorID:   c.AuthorID,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}
