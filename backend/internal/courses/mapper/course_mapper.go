package mapper

import (
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"

	authMapper "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/mapper"
)

// NewCourseResponse creates a CourseResponse from model.Course.
func NewCourseResponse(c *model.Course, includeAuthor bool) *dto.CourseResponse {
	resp := &dto.CourseResponse{
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

	if includeAuthor && c.Author.ID != uuid.Nil {
		resp.Author = authMapper.NewUserResponse(&c.Author)
	}

	return resp
}
