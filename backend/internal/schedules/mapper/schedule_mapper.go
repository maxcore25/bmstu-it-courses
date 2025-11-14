package mapper

import (
	"github.com/google/uuid"

	branchMapper "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/mapper"
	courseMapper "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/mapper"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/model"
)

// NewScheduleResponse creates an ScheduleResponse from model.Schedule, including expanded relations if preloaded.
func NewScheduleResponse(s *model.Schedule) *dto.ScheduleResponse {
	resp := &dto.ScheduleResponse{
		ID:        s.ID,
		CourseID:  s.CourseID,
		BranchID:  s.BranchID,
		StartAt:   s.StartAt,
		EndAt:     s.EndAt,
		Capacity:  s.Capacity,
		Reserved:  s.Reserved,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}

	if s.Course.ID != uuid.Nil {
		resp.Course = courseMapper.NewCourseResponse(&s.Course, false)
	}

	if s.Branch != nil && s.Branch.ID != uuid.Nil {
		resp.Branch = branchMapper.NewBranchResponse(s.Branch)
	}

	return resp
}
