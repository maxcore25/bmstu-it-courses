package mapper

import (
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/model"
)

func NewScheduleResponse(s *model.Schedule) *dto.ScheduleResponse {
	return &dto.ScheduleResponse{
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
}
